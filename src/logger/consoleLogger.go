package logger

import(
	"fmt"
	"sync"
)
type consoleLogger struct{
	messageList chan string
	wait sync.WaitGroup
}

func newconsoleLogger() *consoleLogger {
	consoleLog := &consoleLogger{}
	consoleLog.messageList = make(chan string)
	go func() {
		consoleLog.wait.Add(1)
		defer consoleLog.wait.Done()
		for msg := range consoleLog.messageList{
			fmt.Print(msg)
		}
	}()
	return consoleLog
}
func (consoleLog *consoleLogger) appendMsg(msg string) {
	consoleLog.messageList <- msg
}

func (consoleLog *consoleLogger) exitLogger() {
	consoleLog.appendMsg("Exit Console Logger")
	close(consoleLog.messageList)
	consoleLog.wait.Wait()
}