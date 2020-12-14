package main

import (
	. "MasterServerGo/src/server/sslog"
	"MasterServerGo/src/server"
	_ "MasterServerGo/src/server/serverdb"
)


func main() {
	LogInfo("##########StartApp###########")
	go func(){server.SERVER("0.0.0.0:8088")}()
	WaitForLogger()
}
