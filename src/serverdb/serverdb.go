package serverdb

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "MasterServerGo/src/server/sslog"
	"fmt"
)
//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "test"
)
func init(){
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	_, err := sql.Open("mysql", conn)
	LogInfo("ttt")
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
}