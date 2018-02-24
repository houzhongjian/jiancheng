package service

import (
	"log"
	"net/http"
	"time"

	"github.com/houzhongjian/jiancheng/base"
	"github.com/houzhongjian/jiancheng/db"
	"github.com/houzhongjian/jiancheng/utils"
	"github.com/jinzhu/gorm"
)

//HandleLogin 登录.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		msg := base.User{}
		if err := utils.BindJSON(r, &msg); err != nil {
			log.Printf("%+v\n", err)
			utils.Response(w, "失败")
			return
		}
		user := base.User{}
		if err := db.JcDB.Model(&msg).Where("account = ?", msg.Account).First(&user).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Printf("%+v\n", err)
				utils.Response(w, "账号或者密码错误")
				return
			}
		}
		if utils.MD5(msg.Password) != user.Password {
			log.Println("登录密码错误")
			utils.Response(w, "账号或者密码错误")
			return
		}

		cookie := &http.Cookie{
			Name:     "status",
			Value:    "ok",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour()+1, time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local),
		}
		http.SetCookie(w, cookie)
		//登录成功.
		utils.Response(w, "成功")
		return
	}

	t, err := utils.Display("admin-login")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, nil)
}
