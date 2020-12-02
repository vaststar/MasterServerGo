package handler

import (
	"io"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hetestllo")
}

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/test",testHandler)
}