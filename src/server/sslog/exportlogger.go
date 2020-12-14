package sslog
import(
	"MasterServerGo/src/logger"
	_ "MasterServerGo/src/configure"
)
const(
	LOG_TAG = "SERVER"
)
func LogTrace(msg ...interface{}){
	logger.LOG_TRACE(LOG_TAG, 2, msg...)
}
func LogDebug(msg ...interface{}){
	logger.LOG_DEBUG(LOG_TAG, 2, msg...)
}
func LogInfo(msg ...interface{}){
	logger.LOG_INFO(LOG_TAG, 2, msg...)
}
func LogWarn(msg ...interface{}){
	logger.LOG_WARN(LOG_TAG, 2, msg...)
}
func LogError(msg ...interface{}){
	logger.LOG_ERROR(LOG_TAG, 2, msg...)
}
func WaitForLogger(){
	logger.WaitForAppExit()
}
func init() {
	logger.InitLogger(int(logger.TraceLevel | logger.DebugLevel | logger.InfoLevel | logger.WarnLevel | logger.ErrorLevel))
	logger.AddConsoleLog()
	logger.AddFileLog("./serverlog/server.log", 10, 50*1024*1024)
}