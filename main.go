package main

import (
	"log"

	"github.com/houzhongjian/jiancheng/conf"
	"github.com/houzhongjian/jiancheng/db"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//初始化配置.
	conf.ConfInit()
	log.Println("website port :", conf.Conf.WebsitePost)
	log.Println("db name :", conf.Conf.DbName)
	log.Println("db user :", conf.Conf.DbUser)
	log.Println("db port :", conf.Conf.DbPort)
	log.Println("db pass :", conf.Conf.DbPassword)

	//连接数据库.
	db.DbInit()
	//路由.
	HttpServer()
}
