package handler

import (
	"net/http"
	. "goserver/server/sslog"
	"goserver/server/serverdb"
	"goserver/server/model"
)

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/user",handleIterceptor(userHandler))
	serverMux.HandleFunc("/assets/images",handleIterceptor(imageHandler))
}

func handleIterceptor(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		LogTrace("About to deal request: "+r.RequestURI)
        h(w, r)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	result := serverdb.QueryUser()
	resp := model.Resp{Code:"0", Data:result}
	MarshalJson(w, resp)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	result := serverdb.QuryAllImages()
	resp := model.Resp{Code:"0", Data:result}
	MarshalJson(w, resp)
}
