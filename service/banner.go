package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
	"github.com/jinzhu/gorm"
)

//HandleBanner banner管理.
func HandleBanner(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBanner")
	t, err := utils.Display("admin-banner")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	list := []*base.Banner{}
	err = db.JcDB.Model(&base.Banner{}).Where("is_show = ?", true).Order("sort desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}
	for _, v := range list {
		v.Create = v.CreatedAt.Format(utils.TIMEFORMAT)
		v.Update = v.UpdatedAt.Format(utils.TIMEFORMAT)
	}
	t.Execute(w, map[string]interface{}{"list": list})
}

//HandleBannerAdd 添加广告.
func HandleBannerAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleBannerAdd")
	if r.Method == "POST" {
		msg := &base.Banner{}
		if err := utils.BindJSON(r, &msg); err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "添加失败")
			return
		}

		err := db.JcDB.Model(&base.Banner{}).Create(&msg).Error
		if err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "添加失败")
			return
		}

		utils.Response(w, "添加成功")
		return
	}
	t, err := utils.Display("admin-banner-add")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, nil)
}
