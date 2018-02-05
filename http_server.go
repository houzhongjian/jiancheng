package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/conf"
	"github.com/houzhongjian/jiancheng/service"
)

//HttpServer 注入路由.
func HttpServer() {
	http.HandleFunc("/", service.HandleDefault)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./frontend/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))
	http.HandleFunc("/login", service.HandleLogin)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Conf.WebsitePost), nil); err != nil {
		log.Printf("%+v\n", "端口监听失败")
	}
}
