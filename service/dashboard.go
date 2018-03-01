package service

import (
	"log"
	"net/http"
	"time"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
)

//HandleDashboard 控制台.
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDashboard")
	t, err := utils.Display("admin-index")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	//统计文章数量，总访问量和，今日访问量.
	var articleNum, totalAccess, dayAccess int64
	err = db.JcDB.Model(&base.Article{}).Count(&articleNum).Error
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	err = db.JcDB.Model(&base.Access{}).Count(&totalAccess).Error
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	startT := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	endT := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 59, time.Local)
	err = db.JcDB.Model(&base.Access{}).Where("created_at >= ? and created_at <= ?", startT, endT).Count(&dayAccess).Error
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, map[string]interface{}{"acticle_num": articleNum, "totalAccess": totalAccess, "dayAccess": dayAccess})
}
