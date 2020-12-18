package logger

func AddConsoleLog(level int){
    if logLog != nil{
        logLog.addLoggerInstance(newLogEntry(level, newconsoleLogger()))
    }
}
func AddFileLog(level int, dirPath string, max_days int, maxSingleSize int){
    if logLog != nil{
        logLog.addLoggerInstance(newLogEntry(level, newfileLogger(dirPath, max_days, maxSingleSize)))
    }
}
func WaitForAppExit(){
    if logLog != nil{
        logLog.waitForLoggerExit()
    }
}
func EndLog(){
    if logLog != nil{
        logLog.endLogger()
    }
}
//----------Shouldn't use below as your bare logger-------------
//----------Pls define your own tag and set funcDepth
func LOG_TRACE(tag string, funcDepth int, msg ...interface{}){
    if logLog != nil{
        logLog.logTrace(tag, funcDepth+1, msg...)
    }
}

func LOG_DEBUG(tag string, funcDepth int, msg ...interface{}){
    if logLog != nil{
        logLog.logDebug(tag, funcDepth+1, msg...)
    }
}

func LOG_INFO(tag string, funcDepth int, msg ...interface{}){
    if logLog != nil{
        logLog.logInfo(tag, funcDepth+1, msg...)
    }
}

func LOG_WARN(tag string, funcDepth int, msg ...interface{}){
    if logLog != nil{
        logLog.logWarn(tag, funcDepth+1, msg...)
    }
}

func LOG_ERROR(tag string, funcDepth int, msg ...interface{}){
    if logLog != nil{
        logLog.logError(tag, funcDepth+1, msg...)
    }
}

//loglevel, determine log level 
type logLevel int
const(
    TraceLevel logLevel = 1 << 0
    DebugLevel logLevel = 1 << 1
    InfoLevel  logLevel = 1 << 2
    WarnLevel  logLevel = 1 << 3
    ErrorLevel logLevel = 1 << 4
)
func (p logLevel) String() string {
    switch (p) {
    case TraceLevel:  return "TRACE"
    case DebugLevel:  return "DEBUG"
    case InfoLevel:   return "INFO"
    case WarnLevel:   return "WARN"
    case ErrorLevel:  return "ERROR"
    default:          return "INFO"
    }
}
