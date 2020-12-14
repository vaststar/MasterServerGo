package logger

import(
	"sync"
)

//baseLogger, all output logger should implement this interface
type baseLogger interface{
    writeOneMsg(string)
    exitLogger()
}
type logEntry struct{
	messageList    chan string
	wait           sync.WaitGroup
	outputLogger    baseLogger
}

func newLogEntry(outLogger baseLogger) *logEntry {
	logentry := &logEntry{
		messageList  :  make(chan string),
		outputLogger :  outLogger,
	}
	go func() {
		logentry.wait.Add(1)
		defer logentry.wait.Done()
		for msg := range logentry.messageList{
			logentry.outputLogger.writeOneMsg(msg)
		}
		logentry.outputLogger.exitLogger()
	}()
	return logentry
}

func (logentry *logEntry) appendMsg(msg string) {
	logentry.messageList <- msg
}

func (logentry *logEntry) exitEntry() {
	close(logentry.messageList)
	logentry.wait.Wait()
}
