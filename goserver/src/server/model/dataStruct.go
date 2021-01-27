package model

const (
    SUCCESS                int = 0
	ERROR                  int = -1
	HTTP_NOT_FOUND         int = 404
	HTTP_INVALID_TOKEN     int = 401
	HTTP_ACCESS_FORBIDDEN  int = 403
	SERVER_INTERNAL_ERROR  int = 501
)

type Resp struct {
    Code int      `json:"code"`
    Msg  string      `json:"msg,omitempty"`
    Data interface{} `json:"data,omitempty"`
}

type User struct{
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
}

type Image struct{
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Uri    string    `json:"uri"`
}

type SecretKey struct{
	Id            string    `json:"id"`
	KeySalt       string    `json:"keySalt"`
	ExpireTime    int       `json:"expireTime"`
}