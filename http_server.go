package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/conf"
	"github.com/houzhongjian/jiancheng/filter"
	"github.com/houzhongjian/jiancheng/service"
)

//HttpServer 注入路由.
func HttpServer() {
	http.HandleFunc("/", service.HandleDefault)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./frontend/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./frontend/fonts"))))
	http.Handle("/tables/", http.StripPrefix("/tables/", http.FileServer(http.Dir("./frontend/tables"))))
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("./frontend/lib"))))
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./frontend/upload"))))
	http.HandleFunc("/login", service.HandleLogin)
	http.HandleFunc("/backend/dashboard", filter.Validate(service.HandleDashboard))
	http.HandleFunc("/backend/menu", filter.Validate(service.HandleMenu))
	http.HandleFunc("/backend/menu/add", filter.Validate(service.HandleMenuAdd))
	http.HandleFunc("/backend/menu/edit", filter.Validate(service.HandleMenuEdit))
	http.HandleFunc("/backend/menu/del", filter.Validate(service.HandleMenuDel))
	http.HandleFunc("/backend/article", filter.Validate(service.HandleArticle))
	http.HandleFunc("/backend/article/add", filter.Validate(service.HandleArticleAdd))
	http.HandleFunc("/backend/article/del", filter.Validate(service.HandleArticleDel))
	http.HandleFunc("/backend/banner", filter.Validate(service.HandleBanner))
	http.HandleFunc("/backend/banner/add", filter.Validate(service.HandleBannerAdd))
	http.HandleFunc("/backend/banner/del", filter.Validate(service.HandleBannerDel))
	http.HandleFunc("/upload", filter.Validate(service.HandleUpload))
	http.HandleFunc("/menu", service.HandleDefaultMenu)
	http.HandleFunc("/article/detail", service.HandleDefaultArticleDetail)
	http.HandleFunc("/access", service.HandleDefaultAccess)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Conf.WebsitePost), nil); err != nil {
		log.Printf("%+v\n", "端口监听失败")
	}
}
