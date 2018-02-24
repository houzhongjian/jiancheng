package service

import (
	"log"
	"net/http"

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
	t.Execute(w, nil)
}
