package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
	"github.com/jinzhu/gorm"
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
	for _, v := range list {
		v.Create = v.CreatedAt.Format(utils.TIMEFORMAT)
		v.Update = v.UpdatedAt.Format(utils.TIMEFORMAT)
	}
	t.Execute(w, map[string]interface{}{"list": list})
}

//HandleMenuAdd 添加栏目.
func HandleMenuAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleMenuAdd")
	if r.Method == "POST" {
		menu := base.Menu{}
		if err := utils.BindJSON(r, &menu); err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "添加失败")
			return
		}

		if len(menu.Name) < 1 {
			utils.Response(w, "名称不能为空")
			return
		}

		err := db.JcDB.Model(&menu).Create(&menu).Error
		if err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "添加失败")
			return
		}

		utils.Response(w, "添加成功")
		return
	}
	t, err := utils.Display("admin-menu-add")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, nil)
}

//HandleMenuEdit 栏目编辑.
func HandleMenuEdit(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleMenuEdit")
	if r.Method == "POST" {
		menu := base.Menu{}
		if err := utils.BindJSON(r, &menu); err != nil {
			log.Printf("%+v\n", err)
			return
		}
		log.Println(menu)
		if len(menu.Name) < 1 {
			utils.Response(w, "栏目名称不能为空")
			return
		}
		err := db.JcDB.Model(&base.Menu{}).Where("id = ?", menu.Id).Updates(&menu).Error
		if err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "更新失败")
			return
		}
		utils.Response(w, "更新成功")
		return
	}
	id := r.FormValue("id")
	menu := base.Menu{}
	err := db.JcDB.Model(&menu).Where("id = ?", id).First(&menu).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Printf("%+v\n", err)
			return
		}
		log.Printf("%+v\n", err)
		return
	}
	t, err := utils.Display("admin-menu-edit")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, menu)
}
