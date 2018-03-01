package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const TIMEFORMAT = "2006-01-02 15:04:05"

//Display 分配视图.
func Display(name string) (*template.Template, error) {
	t, err := template.ParseFiles(fmt.Sprintf("./frontend/%s.html", name))
	return t, err
}

//MD5 加密.
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	b := h.Sum(nil)
	str = hex.EncodeToString(b)
	return str
}

//Response 相应请求.
func Response(w http.ResponseWriter, msg interface{}) {
	dist := make(map[string]interface{})
	dist["msg"] = msg
	b, err := json.Marshal(dist)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	w.Write(b)
}

//BindJSON .
func BindJSON(r *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		return err
	}

	log.Println("json:", string(b))
	if err = json.Unmarshal(b, &obj); err != nil {
		log.Printf("%+v\n", err)
		return err
	}
	return nil
}

//Redirect 跳转.
func Redirect(w http.ResponseWriter, path string) {
	w.Write([]byte(fmt.Sprintf("<script>window.location.href='%s'</script>", path)))
}

//FileSuffix 获取文件后缀.
func FileSuffix(str string) string {
	sArr := strings.Split(str, ".")
	return sArr[1]
}

//GetIp 获取ip.
func GetIp(r *http.Request) string {
	ipadd := r.RemoteAddr
	sArr := strings.Split(ipadd, ":")
	return sArr[0]
}
