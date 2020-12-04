package server

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
	"MasterServerGo/src/server/handler"
)

func SERVER() {
	LogInfo("====Start Server====")
	mu := http.NewServeMux()
	handler.InitRouter(mu)

	s := &http.Server{
		Addr:           ":8088",
		Handler:        mu,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	monitorSystem(s)
	s.ListenAndServe()
}

func monitorSystem(server *http.Server){
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	go func() {
		<- quitChan
		server.Close()
	}()
}