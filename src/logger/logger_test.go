package logger

import(
	"testing"
	"MasterServerGo/src/logger"
)

func LogDebugTest(msg ...interface{}){
	logger.LOG_DEBUG("TEST", 2, msg...)
}
func init() {
	logger.InitLogger(int(logger.TraceLevel | logger.DebugLevel | logger.InfoLevel | logger.WarnLevel | logger.ErrorLevel))
	logger.AddConsoleLog()
	logger.AddFileLog("./testlog/test/cute.log", 1, 5*1024*1024)
}
func TestFileLog(t *testing.T){
	t.Log("=====Test Async Log=====")
	defer t.Log("=====Finish Async Log=====")
    quit := make(chan int)
    go func(){
		for i := 0; i < 100; i++{
			LogDebugTest("gg", i)
		}
		quit<-1
	}()
	go func(){
		for i := 0; i < 100; i++{
			t.Log("ss")
			LogDebugTest("tt", i)
		}
		quit<-1
	}()
	sum := 0
	for range quit{
		sum++
		if sum >= 2{
			close(quit)
			logger.EndLog()
			return 
		}
	}
}