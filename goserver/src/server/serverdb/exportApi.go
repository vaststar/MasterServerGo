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

func QueryUser() []User{
	LogDBInfo(" query user")
	var result []User
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
			result = append(result,User{id,name,password})
		}
	}
	err = rows.Err()
	if err != nil {
		LogDBError(err)
	}
	return result
}

func QuryAllImages() []Image{
	LogDBInfo(" query images")
	var result []Image
	if DBDB == nil{
		LogDBError("no db")
		return result
	}
	rows, err := DBDB.db.Query("select * from wedding_images")
	if err != nil {
		LogDBError(err)
		return result
	}
	defer rows.Close()
	var id, name, uri string
	for rows.Next() {
	    err := rows.Scan(&id, &name, &uri)
	    if err != nil {
			LogDBError(err)
	    }else{
			result = append(result, Image{id,name,uri})
		}
	}
	err = rows.Err()
	if err != nil {
		LogDBError(err)
	}
	return result
}