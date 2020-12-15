package main

import (
	"fmt"
	"MasterServerGo/src/configure"
	. "MasterServerGo/src/server/sslog"
	"MasterServerGo/src/server"
	_ "MasterServerGo/src/server/serverdb"
)

func main() {
	//config
	conf, err := configure.ReadConfig("./config/config.json")
	if err != nil{
		panic(err)
	}

	//logger
	if conf.LogConf.ConsoleConf != nil && conf.LogConf.ConsoleConf.Use{
		AddConsoleLog(conf.LogConf.ConsoleConf.LogLevel)
	}
	if conf.LogConf.FileConf != nil && conf.LogConf.FileConf.Use{
		AddFileLog(conf.LogConf.FileConf.LogLevel, conf.LogConf.FileConf.LogPath, conf.LogConf.FileConf.MaxDays, conf.LogConf.FileConf.FileSize)
	}

	LogInfo("##########StartApp###########")
	//init atabase

	//run server
	go func(){
		server.SERVER(fmt.Sprintf("%v:%d",conf.ServerConf.Ip,conf.ServerConf.Port))
	}()
	WaitForLogger()
}
