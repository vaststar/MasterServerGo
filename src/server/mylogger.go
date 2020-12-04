package server
import(
	"MasterServerGo/src/logger"
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