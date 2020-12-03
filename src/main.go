package main

import (
	"net/http"
	"time"
	. "MasterServerGo/src/logger"
	"MasterServerGo/src/handler"
)

func main() {
	LOG_INFO("Starting the application...")
	mu := http.NewServeMux()
	handler.InitRouter(mu)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mu,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}