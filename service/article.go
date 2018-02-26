package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
	"github.com/jinzhu/gorm"
)

//HandleArticle 文章管理.
func HandleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleArticle")
	t, err := utils.Display("admin-article")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	list := []*base.Article{}
	err = db.JcDB.Model(&base.Article{}).Order("id desc").Find(&list).Error
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

//HandleArticleAdd 添加文章.
func HandleArticleAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleArticleAdd")
	if r.Method == "POST" {
		msg := base.Article{}
		if err := utils.BindJSON(r, &msg); err != nil {
			log.Printf("%+v\n", err)
			return
		}

		if len(msg.Title) < 1 {
			utils.Response(w, "标题不能为空")
			return
		}
		if len(msg.Author) < 1 {
			utils.Response(w, "作者不能为空")
			return
		}
		if len(msg.Keyword) < 1 {
			utils.Response(w, "关键字不能为空")
			return
		}
		if len(msg.Description) < 1 {
			utils.Response(w, "描述不能为空")
			return
		}
		if len(msg.Content) < 1 {
			utils.Response(w, "内容不能为空")
			return
		}

		err := db.JcDB.Model(&base.Article{}).Create(&msg).Error
		if err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "添加失败")
			return
		}
		utils.Response(w, "添加成功")
		return
	}
	t, err := utils.Display("admin-article-add")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	menu := []*base.Menu{}
	err = db.JcDB.Model(&base.Menu{}).Where("is_show = ?", true).Find(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("%+v\n", err)
		return
	}

	t.Execute(w, map[string]interface{}{"menu": menu})
}

//HandleArticleDel 删除文章.
func HandleArticleDel(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleArticleDel")
	msg := base.Article{}
	if err := utils.BindJSON(r, &msg); err != nil {
		log.Printf("%+v\n", err)
		utils.Response(w, "删除失败")
		return
	}

	err := db.JcDB.Model(&base.Model{}).Delete(&msg).Error
	if err != nil {
		log.Printf("%+v\n", err)
		utils.Response(w, "删除失败")
		return
	}

	utils.Response(w, "删除成功")
}
