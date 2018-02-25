package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
)

//HandleMenu 栏目管理.
func HandleMenu(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleMenu")
	t, err := utils.Display("admin-menu")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	list := []*base.Menu{}
	err = db.JcDB.Model(&base.Menu{}).Order("sort desc").Find(&list).Error
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, map[string]interface{}{"list": list})
}
