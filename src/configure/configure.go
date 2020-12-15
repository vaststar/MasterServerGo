package configure

import(
	"encoding/json"
	"io/ioutil"
)

//db config
type SqliteConfig struct{
	FilePath     string `json:"filepath"`
	Use          bool   `json:"use"`
}
type MysqlConfig struct{
	UserName 	 string `json:"username"`
	Password 	 string `json:"password"`
	Ip       	 string `json:"ip"`
	Port     	 int    `json:"port"`
	Network  	 string `json:"network"`
	DatabaseName string `json:"database"`
	Use          bool   `json:"use"`
}
type DataBaseConfig struct{
	MysqlConf  *MysqlConfig `json:"mysql,omitempty"`
	SqliteConf *SqliteConfig `json:"sqlite,omitempty"`
}

//logger config
type FileLogConfig struct{
	LogPath     string `json:"path"`
	MaxDays     int    `json:"maxkeepday"`
	FileSize    int    `json:"filesize"`
	LogLevel    int    `json:"level"`
	Use         bool    `json:"use"`
}
type ConsoleLogConfig struct{
	LogLevel    int    `json:"level"`
	Use         bool    `json:"use"`
}
type LogConfig struct{
	FileConf    *FileLogConfig     `json:"fileLog,omitempty"`
	ConsoleConf *ConsoleLogConfig  `json:"consoleLog,omitempty"`
}

//server config
type ServerConfig struct{
	Ip         string  `json:"ip"`
	Port       int     `json:"port"`
}

type Configure struct{
	DbConf  	DataBaseConfig `json:"dbConfig"`
	LogConf 	LogConfig      `json:"logConfig"`
	ServerConf  ServerConfig   `json:"serverConfig"`
}

func ReadConfig(filepath string)(*Configure,error){
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Configure{}
	err = json.Unmarshal(bytes, config)
	return config, err
}

