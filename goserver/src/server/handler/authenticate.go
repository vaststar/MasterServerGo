package handler

import(
	"fmt"
	"time"
	"net/http"
	"strings"
	"goserver/thirdparty/github.com/dgrijalva/jwt-go"
	"goserver/server/model"
	"goserver/server/serverdb"
	. "goserver/server/sslog"
)

type UserClaims struct{
	UserId      string   `json:"userid"`
	jwt.StandardClaims
}

func createToken(userid string, secret model.SecretKey, accessToken bool) (string, error) {
	var expireDuration int
	if accessToken{
		expireDuration = secret.AccessExpireTime;
	} else{
		expireDuration = secret.RefreshExpireTime;
	}
    at := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expireDuration)).Unix(),
		},
    })
    token, err := at.SignedString([]byte(secret.KeySalt))
    if err != nil {
       return "", err
    }
    return token, nil
}

func parseToken(tokenStr string, secret model.SecretKey) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret.KeySalt), nil 
	})
	if err != nil {
		return "", err
	}   

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", fmt.Errorf("Unauthorized token")
	}
}

func requestRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		LogError("Can't parse request form ")
		w.WriteHeader(http.StatusBadRequest)
		MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Error request!"})
	    return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == ""{
		w.WriteHeader(http.StatusBadRequest)
		MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Empty username!"})
		return
	}
	user := serverdb.QueryUserWithName(username)
	if user.Id == ""{
		LogInfo("Can't find user: ",username)
		w.WriteHeader(http.StatusBadRequest)
		MarshalJson(w, model.Resp{Code:model.HTTP_NOT_FOUND, Msg:"Can't find user!"})
		return 
	}
	if password != user.Password{
		LogInfo("Wrong password for user: ",username)
		w.WriteHeader(http.StatusUnauthorized)
		MarshalJson(w, model.Resp{Code:model.HTTP_INVALID_TOKEN,Msg:"Wrong password!"})
		return
	}
	key := serverdb.QueryKeyScretWithUserId(user.Id)
	if key.Id == ""{
		LogInfo("Can't find keys for user: ",username)
		w.WriteHeader(http.StatusInternalServerError)
		MarshalJson(w, model.Resp{Code:model.HTTP_NOT_FOUND, Msg:"No key salt!"})
		return 
	}
	tokenStr,err := createToken(user.Id, key, false)
	if err != nil{
		LogInfo("createToken error for user: ",username)
		w.WriteHeader(http.StatusInternalServerError)
		MarshalJson(w, model.Resp{Code:model.SERVER_INTERNAL_ERROR, Msg:"Create token error!"})
		return
	}
	resp := model.Resp{Code:model.SUCCESS, Data:tokenStr}
	MarshalJson(w, resp)
}

func requestAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		LogError("Can't parse request form ")
		w.WriteHeader(http.StatusBadRequest)
		MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Error request!"})
	    return
	}
	
	refreshToken := r.Form.Get("refreshToken")
	userid := r.Form.Get("userid")
	if refreshToken == "" || userid == ""{
		LogError("Refresh token check failed")
		w.WriteHeader(http.StatusBadRequest)
		MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Empty userid!"})
		return
	}
	key := serverdb.QueryKeyScretWithUserId(userid)
	if key.Id == ""{
		LogInfo("Can't find keys for userid: ",userid)
		w.WriteHeader(http.StatusNotFound)
		MarshalJson(w, model.Resp{Code:model.HTTP_NOT_FOUND, Msg:"No key salt!"})
		return 
	}
	if _, err := parseToken(refreshToken, key); err == nil{
		tokenStr,err := createToken(userid, key, true)
		if err != nil{
			LogInfo("createToken error for user: ",userid)
			w.WriteHeader(http.StatusInternalServerError)
			MarshalJson(w, model.Resp{Code:model.SERVER_INTERNAL_ERROR, Msg:"Create token error!"})
			return
		}
		resp := model.Resp{Code:model.SUCCESS, Data:tokenStr}
		MarshalJson(w, resp)
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Unauthorized access to this resource!"})
	}
}

func validTokenHandlerIterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LogTrace("About to verify request: "+r.RequestURI)
		userid := r.Header.Get("userid")
		if userid == ""{
			w.WriteHeader(http.StatusNotFound)
			MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"No userid!"})
			return 
		}
		key := serverdb.QueryKeyScretWithUserId(userid)
		if key.Id == ""{
			LogInfo("Can't find keys for userid: ",userid)
			w.WriteHeader(http.StatusNotFound)
			MarshalJson(w, model.Resp{Code:model.HTTP_NOT_FOUND, Msg:"No key salt!"})
			return 
		}

		toeknStr := r.Header.Get("Authorization")
		if toeknStr == "" || !strings.HasPrefix(toeknStr, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			MarshalJson(w, model.Resp{Code:model.HTTP_INVALID_TOKEN, Msg:"Token is not valid!"})
			return
		}
		token, err := parseToken(toeknStr[7:],key)
		if err == nil {
			if token ==  userid{
				h(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				MarshalJson(w, model.Resp{Code:model.HTTP_INVALID_TOKEN, Msg:"Token is not valid!"})
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			MarshalJson(w, model.Resp{Code:model.ERROR, Msg:"Unauthorized access to this resource!"})
		}
	}
}