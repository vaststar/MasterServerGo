package handler

import(
	"fmt"
	"time"
	"goserver/thirdparty/github.com/dgrijalva/jwt-go"
	"goserver/server/model"
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