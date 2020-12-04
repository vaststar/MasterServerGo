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

var masterLog *loglog 
func init(){
    masterLog = &loglog{filePathLength:40}
    masterLog.runningLogger()
}

//baseLogger, all output logger should implement this interface
type baseLogger interface{
    appendMsg(string)
    exitLogger()
}

//loglog, truely logger
type loglog struct{
    lock          sync.Mutex
    outLogger     []baseLogger
    level         int 
    wait          sync.WaitGroup
    filePathLength int
}

func (masterLog *loglog) addLoggerInstance(bblog baseLogger){
    if bblog == nil{
        return
    }
    masterLog.lock.Lock()
    defer masterLog.lock.Unlock()
    masterLog.outLogger = append(masterLog.outLogger,bblog)
}
func (masterLog *loglog) writeLog(tag string, level logLevel, skip int, msg ...interface{}){
    if int(level) & masterLog.level != 0{
        logMsg := masterLog.composeLogMessage(tag, level, skip+1, msg)
        masterLog.lock.Lock()
        defer masterLog.lock.Unlock()
        for _, val := range masterLog.outLogger{
            val.appendMsg(logMsg)
        }
    }
}
func (masterLog *loglog) waitForLoggerExit(){
    masterLog.wait.Wait()
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
    signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
    go func() {
        masterLog.wait.Add(1)
        defer masterLog.wait.Done()
        for s := range c {
            switch s {
            case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
                masterLog.logInfo("LOGGER",2,"End Logger")
                masterLog.lock.Lock()
                defer masterLog.lock.Unlock()
                for _, val := range(masterLog.outLogger){
                    val.exitLogger()
                }
                masterLog.outLogger = nil
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
