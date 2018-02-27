package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
	"github.com/jinzhu/gorm"
)

//HandleDefault .
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	t, err := utils.Display("index")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	menu := []*base.Menu{}
	err = db.JcDB.Model(&base.Menu{}).Where("is_show = ?", true).Order("sort desc").Find(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}

	banner := []*base.Banner{}
	err = db.JcDB.Model(&base.Banner{}).Where("is_show = ?", true).Order("sort desc").Find(&banner).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, map[string]interface{}{"menu": menu, "banner": banner})
}
