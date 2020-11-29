package handler

import (
	"io"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hetestllo")
}

func initRouter(serverMux *ServeMux ){
	serverMux.HandleFunc("/test",testHandler)
}