package serverdb

import(
	"database/sql"
	"time"
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
	if dbConn, err := sql.Open(driverName, dataSourceName); err == nil{
		if err = dbConn.Ping(); err == nil{
			svDB := &serverDB{
				db : dbConn,
			}
			LogDBInfo("connect to db success")
			return svDB
		}else{
			LogDBError(err)
		}
	}else{
		LogDBError(err)
	}
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()
	tryCount := 0;
	for{
		select {
		case <- t.C:
			if tryCount > 10{
				LogDBError("try connect db for 5 times and fail")
				return nil
			}
			tryCount++
			LogDBError("start reconnect:",tryCount)
			if dbConn, err := sql.Open(driverName, dataSourceName); err == nil{
				if err = dbConn.Ping(); err == nil{
					svDB := &serverDB{
						db : dbConn,
					}
					LogDBInfo("connect to db success")
					return svDB
				}else{
					LogDBError(err)
				}
			}else{
				LogDBError(err)
			}
		}
	}
}

func (svDB *serverDB)closeDB(){
	if svDB != nil && svDB.db != nil{
		svDB.db.Close()
	}
}

func (svDB *serverDB)executeSQL(){
	
}
