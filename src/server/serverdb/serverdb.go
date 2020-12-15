package serverdb

import(
	"database/sql"
	"sync"
	//. "MasterServerGo/src/server/sslog"
)
var DBDB *ServerDB = &ServerDB{}
var initOnce sync.Once
type ServerDB struct{
	db *sql.DB
}

func InitDB(db *sql.DB){
	initOnce.Do(func(){
		DBDB.db = db
	})
}
