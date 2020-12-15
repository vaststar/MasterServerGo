package logger

import (
    "time"
    "fmt"
    "runtime"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "strconv"
    "bytes"
    "strings"
)

var logLog     *loglog 
var initSync   sync.Once
func init(){
    logLog = &loglog{
        filePathLength : 40,
        isFinished     : false,
    }
    logLog.wait.Add(1)
    logLog.runningLogger()
}
//loglog, truely logger
type loglog struct{
    lock           sync.Mutex
    logEntry       []*logEntry//base log entry, eg. console, filelog
    level          int 
    wait           sync.WaitGroup
    filePathLength int
    isFinished     bool

    waitForEndSync sync.Once
    endLogSync     sync.Once
}
func (masterLog *loglog) addLoggerInstance(logentry *logEntry){
    if logentry == nil{
        return
    }
    masterLog.lock.Lock()
    defer masterLog.lock.Unlock()
    if masterLog.isFinished {
        return
    }
    masterLog.level |= logentry.level
    masterLog.logEntry = append(masterLog.logEntry,logentry)
}
func (masterLog *loglog) writeLog(tag string, level logLevel, skip int, msg ...interface{}){
    masterLog.lock.Lock()
    sholdAppendLog := int(level) & masterLog.level != 0
    if !sholdAppendLog || masterLog.isFinished{
        masterLog.lock.Unlock()
        return
    }
    masterLog.lock.Unlock()
    //compose log and append to entries
    logMsg := masterLog.composeLogMessage(tag, level, skip+1, msg)
    masterLog.lock.Lock()
    defer masterLog.lock.Unlock()
    for _, val := range masterLog.logEntry{
        val.appendMsg(level, logMsg)
    }
}
func (masterLog *loglog) waitForLoggerExit(){
    masterLog.waitForEndSync.Do(func(){masterLog.wait.Wait()})
}
func (masterLog *loglog) endLogger(){
    masterLog.endLogSync.Do(func(){
        masterLog.cleanUpLogger()
        masterLog.wait.Wait()
    })
}
func (masterLog *loglog) cleanUpLogger(){
    masterLog.lock.Lock()
    defer masterLog.lock.Unlock()
    if masterLog.isFinished {
        return
    }
    masterLog.isFinished = true
    for _, val := range(masterLog.logEntry){
        val.exitEntry()
    }
    masterLog.logEntry = nil
    masterLog.wait.Done()
}

func (masterLog *loglog) logTrace(tag string, skip int, msg ...interface{}){
    masterLog.writeLog(tag, TraceLevel, skip+1, msg...)
}
func (masterLog *loglog) logDebug(tag string, skip int, msg ...interface{}){
    masterLog.writeLog(tag, DebugLevel, skip+1, msg...)
}
func (masterLog *loglog) logInfo(tag string, skip int, msg ...interface{}){
    masterLog.writeLog(tag, InfoLevel, skip+1, msg...)
}
func (masterLog *loglog) logWarn(tag string, skip int, msg ...interface{}){
    masterLog.writeLog(tag, WarnLevel, skip+1, msg...)
}
func (masterLog *loglog) logError(tag string, skip int, msg ...interface{}){
    masterLog.writeLog(tag, ErrorLevel, skip+1, msg...)
}
//should call this while ending program
func (masterLog *loglog) runningLogger(){
    c := make(chan os.Signal)
    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    go func() {
        for s := range c {
            switch s {
            case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
                masterLog.logInfo("LOGGER",2,"End Logger")
                masterLog.cleanUpLogger()
                return
            }
        }
    }()
}

//log util
func (masterLog *loglog) composeLogMessage(tag string, level logLevel, skip int, msg ...interface{}) string{
    log_time  := time.Now().UTC().Format("2006-01-02T15:04:05.000")
    log_level := level.String()
    log_pid   := goID()
    pc , file, line, _ := runtime.Caller(skip)
    file = file + "(" + fmt.Sprint(line) + ")"
    if len(file) > masterLog.filePathLength {
        file = file[len(file)-masterLog.filePathLength:]
    }else if len(file) < masterLog.filePathLength {
        file = fmt.Sprintf(fmt.Sprintf("\\%%ds",masterLog.filePathLength), file)
    }
    funcName := runtime.FuncForPC(pc).Name()
    if split_list := strings.Split(funcName,"/"); len(split_list) > 0{
        funcName = split_list[len(split_list) -1]
    }
    return log_time + " [" + log_level + "] ["  + file + "] [" + 
           fmt.Sprint(log_pid)+ "] [" + tag + "] [" + funcName + "] - " +fmt.Sprintln(msg...)
}

func goID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    id, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
