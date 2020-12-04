package handler

import (
	"io"
	"net/http"
	. "MasterServerGo/src/logger"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	LOG_TRACE("test",1,"hello log")
	io.WriteString(w, "hetestllo")
}

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/test",testHandler)
}