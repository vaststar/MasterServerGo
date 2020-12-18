package serverdb

import(
	"database/sql"
	. "goserver/server/sslog"
	_ "goserver/thirdparty/github.com/mattn/go-sqlite3"
	_ "goserver/thirdparty/github.com/go-sql-driver/mysql"
)
type serverDB struct{
	db *sql.DB
}

var supportedDriver map[string]bool = map[string]bool{"sqlite3":true,"mysql":true}

func createDB(driverName string, dataSourceName string) *serverDB{
	if !supportedDriver[driverName]{
		LogDBError(driverName," not supported.")
		return nil
	}
	dbConn, err := sql.Open(driverName, dataSourceName)
	if err != nil{
		LogDBError(err)
		return nil
	}    
	if err := dbConn.Ping(); err != nil {
		LogDBError(err)
		return nil
	}
	svDB := &serverDB{
		db : dbConn,
	}
	return svDB
}

func (svDB *serverDB)closeDB(){
	if svDB != nil && svDB.db != nil{
		svDB.db.Close()
	}
}

func (svDB *serverDB)executeSQL(){
	
}
