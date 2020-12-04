package logger

import(
	"sync"
	"path/filepath"
	"strings"
	"os"
	"time"
	"sort"
	"regexp"
	"strconv"
)
type fileLogger struct{
	messageList 	 chan string
	wait        	 sync.WaitGroup
	dirPath      	 string  
	baseFileName     string
	singleFileSize   int
	maxKeepDays      int
	mCurrentSize     int
	logFile          *os.File
}

func newfileLogger(file_path string, max_days int, max_singleSize int) *fileLogger {
	absPath,_ := filepath.Abs(file_path)
	fileLog := &fileLogger{
		maxKeepDays      :  max_days,
		singleFileSize   :  max_singleSize,
		dirPath          :  filepath.Dir(absPath),
		baseFileName     :  strings.TrimSuffix(filepath.Base(absPath), filepath.Ext(absPath)),
	}
    os.MkdirAll(fileLog.dirPath, os.ModePerm);
	fileLog.readyForLog(0)
	
	fileLog.messageList = make(chan string)
	go func() {
		fileLog.wait.Add(1)
		defer fileLog.wait.Done()
		for msg := range fileLog.messageList{
			fileLog.writeOneLog(msg)
		}
		fileLog.closeCurrentFile()
	}()
	return fileLog
}
func (fileLog *fileLogger) appendMsg(msg string) {
	fileLog.messageList <- msg
}

func (fileLog *fileLogger) exitLogger() {
	close(fileLog.messageList)
	fileLog.wait.Wait()
}

func (fileLog *fileLogger) writeOneLog(msg string) {
	if fileLog.readyForLog(len(msg)){
		if _, err := fileLog.logFile.WriteString(msg); err == nil{
			fileLog.mCurrentSize += len(msg)
		}
	}
}

func (fileLog *fileLogger) readyForLog(messageSize int) bool{
	fileLog.doRollOver(messageSize)
	fileLog.removeOldFile()

	if fileLog.logFile == nil {
		currentFilePath := filepath.Join(fileLog.dirPath, fileLog.baseFileName +"-"+time.Now().Format("2006-01-02")+".log")
		file, err := os.OpenFile(currentFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return false
		}
		fileinfo, err := os.Stat(currentFilePath)
		if err != nil {
			return false
		}
		fileLog.logFile = file
		fileLog.mCurrentSize = int(fileinfo.Size())
		return true
	}
	return true
}

func (fileLog *fileLogger) doRollOver(messageSize int) bool{
	if messageSize > fileLog.singleFileSize {
		return false
	}
	if fileLog.mCurrentSize + messageSize <= fileLog.singleFileSize {
		return true
	}

	fileLog.closeCurrentFile()
	//rename all file
	files,_ := filepath.Glob(fileLog.baseFileName+"-"+time.Now().Format("2006-01-02")+"\\.log.*?")
	sort.Slice(files, func(i, j int) bool{
		if r := regexp.MustCompile(".*?\\.log$"); r.MatchString(files[i]){
			return false
		}
		if r := regexp.MustCompile(".*?\\.log$"); r.MatchString(files[j]){
			return true
		}
		r := regexp.MustCompile(".*?\\.log\\.(\\d+)$")
		if r.MatchString(files[i]) && r.MatchString(files[j]) {
			numi_list := r.FindStringSubmatch(files[i])
			numi,_ := strconv.Atoi(numi_list[1])
			numj_list := r.FindStringSubmatch(files[j])
			numj,_ := strconv.Atoi(numj_list[1])
			return numi > numj
		}
		return true
	})
	for _, itemFile := range files{
		if r := regexp.MustCompile(".*?\\.log$"); r.MatchString(itemFile) {
			os.Rename(itemFile,itemFile+".1")
		} else if r := regexp.MustCompile("(.*?\\.log\\.)(\\d+)$"); r.MatchString(itemFile){
			params := r.FindStringSubmatch(itemFile)
			old_num,_ := strconv.Atoi(params[2])
			newFileName := params[1] + strconv.Itoa(old_num+1)
			os.Rename(itemFile, newFileName)
		}
	}
	return true
}

func (fileLog *fileLogger) closeCurrentFile() {
	if fileLog.logFile != nil {
		fileLog.logFile.Close()
		fileLog.logFile = nil
		fileLog.mCurrentSize = 0
	}
}

func (fileLog *fileLogger) removeOldFile() {
	files,_ := filepath.Glob(fileLog.baseFileName+"-(\\d{4}-\\d{2}-\\d{2})\\.log"+".*?")
	r := regexp.MustCompile(fileLog.baseFileName+"-(\\d{4}-\\d{2}-\\d{2})\\.log"+".*?")
	for _, itemFile := range files {
		params := r.FindStringSubmatch(itemFile)
		fileTime,_ := time.Parse("2006-01-02",params[1])
		if time.Now().Sub(fileTime) > time.Duration(fileLog.maxKeepDays*24)*time.Hour{
			os.Remove(itemFile)
		}
	}
}