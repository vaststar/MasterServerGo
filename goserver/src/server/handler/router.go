package handler

import (
	"io"
	"net/http"
	. "goserver/server/sslog"
	"goserver/server/serverdb"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	LogInfo("hello log")
	io.WriteString(w, serverdb.QueryUser())
}

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/test",testHandler)
}