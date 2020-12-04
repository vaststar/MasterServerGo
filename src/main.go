package main

import (
	//"path/filepath"
	"MasterServerGo/src/logger"
	"MasterServerGo/src/server"
)

func init() {
	logger.InitLogger(int(logger.TraceLevel | logger.DebugLevel | logger.InfoLevel | logger.WarnLevel | logger.ErrorLevel))
	logger.AddConsoleLog()
	logger.AddFileLog("./testlog/cute.log", 10, 50*1024*1024)
}
func main() {
	logger.LOG_INFO("Main",1,"##########StartApp###########")
	go func(){server.SERVER()}()
	logger.WaitForLogger()
}
