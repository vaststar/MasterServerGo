package configure

import(
	"encoding/json"
	"fmt"
)

type SqliteConfig struct{
	FilePath     string `json:"filepath"`
}
type MysqlConfig struct{
	UserName 	 string `json:"username"`
	Password 	 string `json:"password"`
	Ip       	 string `json:"ip"`
	Port     	 int    `json:"port"`
	Network  	 string `json:"network"`
	DatabaseName string `json:"database"`
}
type DataBaseConfig struct{
	MysqlConf  *MysqlConfig `json:"mysql,omitempty"`
	SqliteConf *SqliteConfig `json:"sqlite,omitempty"`
}
type LogConfig struct{

}
type Configure struct{
	DbConfig DataBaseConfig `json:"database"`
}

func init(){
	testStr := `{"database":{"sqlite":{"filepath":"ggggg"}}}`
	a := new(Configure)
	err := json.Unmarshal([]byte(testStr), a)
    if err != nil {
        fmt.Print(err)
	} 
	fmt.Print(a.DbConfig.SqliteConf.FilePath)
}

