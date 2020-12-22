package logger

import(
	"path/filepath"
	"strings"
	"os"
	"time"
	"sort"
	"regexp"
	"strconv"
	"fmt"
)
type fileLogger struct{
	dirPath      	 string  
	baseFileName     string
	singleFileSize   int
	maxKeepDays      int
	mCurrentSize     int
	mCurrentDate     string
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
	return fileLog
}

func (fileLog *fileLogger) exitLogger() {
	fileLog.closeCurrentFile()
}

func (fileLog *fileLogger) writeOneMsg(msg string) {
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
		currentFilePath := fileLog.getRequiredFilePath()
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
		fmt.Println("[FileLogger error]message too large, max size is: ",fileLog.singleFileSize)
		return false
	}
	if fileLog.mCurrentDate != time.Now().Format("2006-01-02"){
		fileLog.closeCurrentFile()
		fileLog.mCurrentDate = time.Now().Format("2006-01-02")
		return true
	}
	if fileLog.mCurrentSize + messageSize <= fileLog.singleFileSize {
		return true
	}

	fileLog.closeCurrentFile()
	//rename todays file
	files,_ := filepath.Glob(fileLog.dirPath+"/*")
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
		if r := regexp.MustCompile(".*?"+fileLog.mCurrentDate+"\\.log$"); r.MatchString(itemFile) {
			os.Rename(itemFile,itemFile+".1")
		} else if r := regexp.MustCompile("(.*?"+fileLog.mCurrentDate+"\\.log\\.)(\\d+)$"); r.MatchString(itemFile){
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
	files,_ := filepath.Glob(fileLog.dirPath+"/*")
	r := regexp.MustCompile(fileLog.baseFileName+"-(\\d{4}-\\d{2}-\\d{2})\\.log"+".*?")
	for _, itemFile := range files {
		if r.MatchString(itemFile) {
			params := r.FindStringSubmatch(itemFile)
			fileTime,_ := time.Parse("2006-01-02",params[1])
			if time.Now().Sub(fileTime) > time.Duration(fileLog.maxKeepDays*24)*time.Hour{
				os.Remove(itemFile)
			}
		}
	}
}

func (fileLog *fileLogger) getRequiredFilePath() string{
	return filepath.Join(fileLog.dirPath, fileLog.baseFileName +"-"+fileLog.mCurrentDate+".log")
}