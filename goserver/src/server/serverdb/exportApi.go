package serverdb

import(
	"sync"
	"os"
	"bufio"
	"strings"
	. "goserver/server/sslog"
	"goserver/server/model"
)

var DBDB *serverDB
var initOnce sync.Once
var closeOnce sync.Once

func InitDB(driverName string, dataSourceName string){
	initOnce.Do(func(){
		LogDBInfo("====Init DataBase", driverName, dataSourceName, "====")
		DBDB = createDB(driverName, dataSourceName)
	})
}

func CloseDB(){
	closeOnce.Do(func(){
		if DBDB != nil{
			DBDB.closeDB()
		}
	})
}

func ExecuteFiles(files []string){
	if DBDB == nil || len(files) == 0{
		return
	}
	for _, val := range files{
		if file, err := os.Open(val); err == nil{
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines) 
			var text string
			for scanner.Scan() { 
				line := strings.TrimSuffix(strings.TrimSpace(scanner.Text()),"\n")
				if strings.HasPrefix(line, "-") || len(line) == 0{
					continue
				}else if strings.HasSuffix(line,";"){
					text += line
					DBDB.db.Exec(text)
					text = ""
				}else{
					text += line
				}
			} 
		}else {
			LogDBError(val, " can't open, error: ", err)
		}
	}
}

func QueryUser() []model.User{
	var result []model.User
	if DBDB == nil{
		LogDBError("no db")
		return result
	}
	rows, err := DBDB.db.Query("select * from identity")
	if err != nil {
		LogDBError(err)
		return result
	}
	defer rows.Close()
	var id, name, password string
	for rows.Next() {
	    err := rows.Scan(&id, &name, &password)
	    if err != nil {
			LogDBError(err)
		}else{
			result = append(result,model.User{id,name,password})
		}
	}
	err = rows.Err()
	if err != nil {
		LogDBError(err)
	}
	return result
}

func QuryKeyScrets() []model.SecretKey{
	var result []model.SecretKey
	if DBDB == nil{
		LogDBError("no db")
		return result
	}
	rows, err := DBDB.db.Query("select * from secret_key")
	if err != nil {
		LogDBError(err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var tempVal model.SecretKey
	    err := rows.Scan(&tempVal.Id, &tempVal.KeySalt, &tempVal.ExpireTime)
	    if err != nil {
			LogDBError(err)
	    }else{
			result = append(result, tempVal)
		}
	}
	err = rows.Err()
	if err != nil {
		LogDBError(err)
	}
	return result
}