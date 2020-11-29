package main

import (
	"io"
	"net/http"
	"time"
	"log"
	//"server/handler"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hetestllo")
}

func initRouter(serverMux *Server  ){
	serverMux.HandleFunc("/test",testHandler)
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/echo",echoHandler)
	initRouter(&mu)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mu,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}