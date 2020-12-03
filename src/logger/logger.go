package logger

// import (
//     "log"
//     "os"
// )
var masterLog *loglog//main log instance, should be init in main
func LOG_DEBUG(msg ...interface{}){

}

func LOG_INFO(msg ...interface{}){

}

func LOG_WARN(msg ...interface{}){

}

func LOG_ERROR(msg ...interface{}){

}

type logLevel int32
const(
    traceLevel logLevel = 0
    debugLevel logLevel = 1
    infoLevel logLevel = 2
    warnLevel logLevel = 3
    errorLevel logLevel = 4
)

type loglog struct{
    outLogger []*baseLogger

}

func (masterLog *loglog) writeLog(){

}

type baseLogger interface{
    Append(string)
}