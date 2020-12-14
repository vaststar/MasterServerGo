package logger

import(
	"fmt"
)
type consoleLogger struct{
}

func newconsoleLogger() *consoleLogger {
	consoleLog := &consoleLogger{}
	return consoleLog
}
func (consoleLog *consoleLogger) writeOneMsg(msg string) {
	fmt.Print(msg)
}
func (consoleLog *consoleLogger) exitLogger() {
}