package handler

import(
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
    . "goserver/server/sslog"
)
// MarshalJson 把对象以json格式放到response中
func MarshalJson(w http.ResponseWriter, v interface{}) error {
    data, err := json.Marshal(v)
    if err != nil {
        return err
	}
	LogTrace(string(data))
    w.Write(data)
    return nil
}

// UnMarshalJson 从request中取出对象
func UnMarshalJson(req *http.Request, v interface{}) error {
    result, err := ioutil.ReadAll(req.Body)
    if err != nil {
		LogError(req.RequestURI+" request parse error")
        return err
    }
    json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
    return nil
}