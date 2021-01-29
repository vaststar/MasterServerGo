package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	. "goserver/server/sslog"
	"goserver/server/serverdb"
	"goserver/server/model"
	"goserver/server/configure"
)

func InitRouter(serverMux *http.ServeMux ){
	serverMux.HandleFunc("/authenticate/requestAccessToken",requestAccessTokenHandler)
	serverMux.HandleFunc("/authenticate/requestRefreshToken",requestRefreshTokenHandler)
	serverMux.HandleFunc("/users",validTokenHandlerIterceptor(userHandler))
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
	resp := model.Resp{Code:model.SUCCESS, Data:result}
	MarshalJson(w, resp)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	files,_ := filepath.Glob(configure.GetConfig().AssetsConf.ImagesPath+"*")
	var result []model.Image
	for index, str := range files{
		result = append(result, model.Image{strconv.Itoa(index),filepath.Base(str),configure.GetConfig().AssetsConf.ImagesUri + filepath.Base(str)})
	}
	resp := model.Resp{Code:model.SUCCESS, Data:result}
	MarshalJson(w, resp)
}
