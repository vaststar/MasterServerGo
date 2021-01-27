package handler

import(
	"fmt"
	"time"
	"net/http"
	"goserver/thirdparty/github.com/dgrijalva/jwt-go"
	"goserver/server/model"
	"goserver/server/serverdb"
	. "goserver/server/sslog"
)

type UserClaims struct{
	UserId      string   `json:"userid"`
	jwt.StandardClaims
}

func CreateToken(userid string, secret model.SecretKey) (string, error) {
    at := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(secret.ExpireTime)).Unix(),
		},
    })
    token, err := at.SignedString([]byte(secret.KeySalt))
    if err != nil {
       return "", err
    }
    return token, nil
}

func ParseToken(tokenStr string, secret model.SecretKey) (string, error) {
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

func requestTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		LogError("Can't parse request form ")
	    return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == ""{
		LogError("Empty username ")
		return
	}
	
}