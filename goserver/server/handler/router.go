package handler

import (
	"io"
	"net/http"
	. "goserver/server/sslog"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	LogInfo("hello log")
	io.WriteString(w, "hetestllo")
}

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/test",testHandler)
}