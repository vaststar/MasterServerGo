package serverdb

type Error struct{
	ErrorCode string `json:"errorCode"`
	ErrorString string `json:"errorString"`
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