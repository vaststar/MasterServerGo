package main

import (
	//"path/filepath"
	. "MasterServerGo/src/server/sslog"
	"MasterServerGo/src/server"
	//_ "MasterServerGo/src/serverdb"
)


func main() {
	LogInfo("##########StartApp###########")
	go func(){server.SERVER("0.0.0.0:8088")}()
	WaitForLogger()
}
