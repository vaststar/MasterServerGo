package sslog
import(
	"goserver/thirdparty/logger"
	_ "MasterServerGo/src/configure"
)
const(
	LOG_SERVER_TAG = "SERVER"
	LOG_DB_TAG = "DB"
)
//server or main log
func LogTrace(msg ...interface{}){
	logger.LOG_TRACE(LOG_SERVER_TAG, 2, msg...)
}
func LogDebug(msg ...interface{}){
	logger.LOG_DEBUG(LOG_SERVER_TAG, 2, msg...)
}
func LogInfo(msg ...interface{}){
	logger.LOG_INFO(LOG_SERVER_TAG, 2, msg...)
}
func LogWarn(msg ...interface{}){
	logger.LOG_WARN(LOG_SERVER_TAG, 2, msg...)
}
func LogError(msg ...interface{}){
	logger.LOG_ERROR(LOG_SERVER_TAG, 2, msg...)
}
// db log
func LogDBTrace(msg ...interface{}){
	logger.LOG_TRACE(LOG_DB_TAG, 2, msg...)
}
func LogDBDebug(msg ...interface{}){
	logger.LOG_DEBUG(LOG_DB_TAG, 2, msg...)
}
func LogDBInfo(msg ...interface{}){
	logger.LOG_INFO(LOG_DB_TAG, 2, msg...)
}
func LogDBWarn(msg ...interface{}){
	logger.LOG_WARN(LOG_DB_TAG, 2, msg...)
}
func LogDBError(msg ...interface{}){
	logger.LOG_ERROR(LOG_DB_TAG, 2, msg...)
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