package service

import (
	"log"
	"net/http"

	"github.com/houzhongjian/jiancheng/utils"
)

//HandleDefault .
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	t, err := utils.Display("index")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, nil)
}
