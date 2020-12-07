package main

import (
	//"path/filepath"
	. "MasterServerGo/src/server/sslog"
	"MasterServerGo/src/server"
)

func main() {
	LogInfo("##########StartApp###########")
	go func(){server.SERVER()}()
	WaitForLogger()
}
