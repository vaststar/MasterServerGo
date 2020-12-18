package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"goserver/server/configure"
	. "goserver/server/sslog"
	"goserver/server/serverdb"
	"goserver/server"
)
var configPath = flag.String("config", "./config/config.json", "Input Configure FilePath")

func main() {
	flag.Parse()
	//config
	conf, err := configure.ReadConfig(*configPath)
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
	var driverName, dataSourceName string
	var files []string
	if conf.DbConf.SqliteConf != nil && conf.DbConf.SqliteConf.Use{
		driverName, dataSourceName, files = "sqlite3", conf.DbConf.SqliteConf.DbPath, conf.DbConf.SqliteConf.SqlFiles
		//mkdir
		os.MkdirAll(filepath.Dir(dataSourceName), os.ModePerm);
		dataSourceName = "file:"+dataSourceName+"?cache=shared"
	} else if conf.DbConf.MysqlConf != nil && conf.DbConf.MysqlConf.Use{
		driverName = "mysql"
		dataSourceName = fmt.Sprintf("%s:%s@%s(%s:%d)/%s",conf.DbConf.MysqlConf.UserName, conf.DbConf.MysqlConf.Password, 
												conf.DbConf.MysqlConf.Protocol, conf.DbConf.MysqlConf.Ip, 
												conf.DbConf.MysqlConf.Port, conf.DbConf.MysqlConf.DbName)
		files = conf.DbConf.MysqlConf.SqlFiles
	}
	if len(driverName) > 0{
		serverdb.InitDB(driverName, dataSourceName)
		serverdb.ExecuteFiles(files)
		defer serverdb.CloseDB()
	} else {
		LogError("Missing sql configure.")
	}

	//run server
	go func(){
		server.SERVER(fmt.Sprintf("%v:%d",conf.ServerConf.Ip,conf.ServerConf.Port))
	}()
	WaitForLogger()
}
