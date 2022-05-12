package server

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
	"goserver/server/handler"
	. "goserver/server/sslog"
)

func SERVER(addr string) {
	LogInfo("====Start Server On: ",addr,"====")
	mu := http.NewServeMux()
	handler.InitRouter(mu)

	s := &http.Server{
		Addr:           addr,
		Handler:        mu,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	monitorSystem(s)
	s.ListenAndServe()
	LogInfo("====Finish Server On: ",addr,"====")
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