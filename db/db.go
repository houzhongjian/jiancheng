package db

import (
	"fmt"
	"log"

	"github.com/houzhongjian/jiancheng/base"
	cf "github.com/houzhongjian/jiancheng/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

//DbInit 初始化数据库.
func DbInit() {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cf.Conf.DbHost, cf.Conf.DbPort, cf.Conf.DbUser, cf.Conf.DbPassword, cf.Conf.DbName, "disable")
	db, err := gorm.Open(cf.Conf.DbUser, args)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	db.SingularTable(true)

	Db = db
}

func AutoMig() {
	Db.AutoMigrate(
		&base.User{},
	)
}
