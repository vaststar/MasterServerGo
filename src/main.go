package main

import (
	"fmt"
	"database/sql"
	"MasterServerGo/src/configure"
	. "MasterServerGo/src/server/sslog"
	"MasterServerGo/src/server"
	"MasterServerGo/src/server/serverdb"
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
	defer WaitForLogger()

	LogInfo("##########StartApp###########")
	//init atabase
	if conf.DbConf.SqliteConf != nil && conf.DbConf.SqliteConf.Use{
		db, err := sql.Open("sqlite", conf.DbConf.SqliteConf.DbPath)
		if err != nil {
			LogError(err)
			return
		}
		serverdb.InitDB(db)
	} else if conf.DbConf.MysqlConf != nil && conf.DbConf.MysqlConf.Use{
		str := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",conf.DbConf.MysqlConf.UserName, conf.DbConf.MysqlConf.Password, 
												conf.DbConf.MysqlConf.Network, conf.DbConf.MysqlConf.Ip, 
												conf.DbConf.MysqlConf.Port, conf.DbConf.MysqlConf.DatabaseName)
		db, err := sql.Open("mysql", str)
		if err != nil {
			LogError(err)
			return
		}
		serverdb.InitDB(db)
	}

	//run server
	go func(){
		server.SERVER(fmt.Sprintf("%v:%d",conf.ServerConf.Ip,conf.ServerConf.Port))
	}()
}
