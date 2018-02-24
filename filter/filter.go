package filter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/utils"
)

// Validate 过滤掉没有登录的请求.
func Validate(handle http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("status")
		if err != nil {
			log.Printf("%+v\n", "未登录")
			utils.Redirect(w, "/login")
			return
		}
		if fmt.Sprintf("%s", cookie.Value) != "ok" {
			log.Printf("%+v\n", "未登录")
			utils.Redirect(w, "/login")
			return
		}
		handle(w, r)
	})
}
