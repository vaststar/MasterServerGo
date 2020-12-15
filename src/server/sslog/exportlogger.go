package sslog
import(
	"MasterServerGo/thirdparty/logger"
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
func AddConsoleLog(level int){
	logger.AddConsoleLog(level)
}
func AddFileLog(level int, dirPath string, max_days int, maxSingleSize int){
	logger.AddFileLog(level, dirPath, max_days, maxSingleSize)
}
func WaitForLogger(){
	logger.WaitForAppExit()
}