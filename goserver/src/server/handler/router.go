package handler

import (
	"io"
	"net/http"
	. "goserver/server/sslog"
	"goserver/server/serverdb"
	"encoding/json"
)

func checkQueryError(res []byte, err error)string{
	var result string
	if err !=nil {
		temp,_ := json.Marshal(serverdb.Error{ErrorCode:"0",ErrorString:"parse result fail"})
		result = string(temp)
	}else{
		result = string(res)
	}
	return result
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	LogInfo("hello userHandler")
	result := checkQueryError(json.Marshal(serverdb.QueryUser()))
	io.WriteString(w, result)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	LogInfo("hello imageHandler")
	result := checkQueryError( json.Marshal(serverdb.QuryAllImages()) )
	io.WriteString(w, result)
}

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/user",userHandler)
	serverMux.HandleFunc("/assets/images",imageHandler)
}