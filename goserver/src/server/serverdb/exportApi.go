package serverdb

import(
	"sync"
	"os"
	"bufio"
	"strings"
	. "goserver/server/sslog"
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

func QueryUser() string{
	LogDBInfo(" query user")
	if DBDB == nil{
		return "no db"
	}
	rows, err := DBDB.db.Query("select * from identity")
	if err != nil {
		LogDBError(err)
	}
	defer rows.Close()
	var id, name, password string
	var result string
	for rows.Next() {
	    err := rows.Scan(&id, &name, &password)
	    if err != nil {
			LogDBError(err)
	    }
		LogDBInfo("one done ")
	    result += "id: " +id + ", name: "+ name+ ", password: "+ password
	}
	err = rows.Err()
	if err != nil {
		LogDBError(err)
	}
	return result
}