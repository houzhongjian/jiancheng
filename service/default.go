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

//HandleDefaultMenu 首页栏目.
func HandleDefaultMenu(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDefaultMenu")
	t, err := utils.Display("menu")
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

	menuId := r.FormValue("id")
	menuNow := &base.Menu{}
	err = db.JcDB.Model(&base.Menu{}).Where("id = ?", menuId).First(&menuNow).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}

	article := []*base.Article{}
	err = db.JcDB.Model(&base.Article{}).Where("menu_id = ?", menuId).Order("id desc").Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}
	for _, v := range article {
		v.Create = v.CreatedAt.Format(utils.TIMEFORMAT)
		v.Update = v.UpdatedAt.Format(utils.TIMEFORMAT)
	}
	t.Execute(w, map[string]interface{}{"menu": menu, "banner": banner, "menunow": menuNow, "article": article})
}

//HandleDefaultArticleDetail 文章详细.
func HandleDefaultArticleDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDefaultArticleDetail")
	id := r.FormValue("id")
	article := base.QueryArticle{}
	err := db.JcDB.Model(&base.Article{}).Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
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

	t, err := utils.Display("detail")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	articleId := r.FormValue("id")
	arti := base.Article{}
	err = db.JcDB.Model(&base.Article{}).Select("menu_id").Where("id = ?", articleId).First(&arti).Error
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	menuNow := &base.Menu{}
	err = db.JcDB.Model(&base.Menu{}).Where("id = ?", arti.MenuId).First(&menuNow).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}

	t.Execute(w, map[string]interface{}{"detail": article, "banner": banner, "menu": menu, "menunow": menuNow})
}
