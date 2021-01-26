package model

type Resp struct {
    Code string      `json:"code"`
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