package logger

import(
	"testing"
)

func LogDebugTest(msg ...interface{}){
	LOG_DEBUG("TEST", 2, msg...)
}
func init() {
	allLevel := int(TraceLevel | DebugLevel | InfoLevel | WarnLevel | ErrorLevel)
	AddConsoleLog(allLevel)
	AddFileLog(allLevel, "./testlog/test/cute.log", 1, 5*1024*1024)
}
func TestFileLog(t *testing.T){
	t.Log("=====Test Async Log=====")
	defer t.Log("=====Finish Async Log=====")
	n := 3//chan number
	quit := make(chan int,n)
	for u:=0;u<n;u++{
		go func(num int){
			for i := 0; i < 10000; i++{
				LogDebugTest(num,"<<", i)
			}
			quit<-1
		}(u)
	}
	sum := 0
	for range quit{
		sum++
		if sum >= n{
			close(quit)
			EndLog()
			return 
		}
	}
}