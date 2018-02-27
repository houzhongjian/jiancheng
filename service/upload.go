package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/houzhongjian/jiancheng/utils"
)

//HandleUpload 上传文件.
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleUpload")
	file, handle, err := r.FormFile("pic")
	if err != nil {
		log.Printf("%+v\n", err)
		utils.Response(w, "上传成功")
		return
	}
	filename := fmt.Sprintf("%d.%s", time.Now().Unix(), utils.FileSuffix(handle.Filename))
	f, err := os.OpenFile(fmt.Sprintf("./frontend/upload/%s", filename), os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	if err != nil {
		log.Printf("%+v\n", err)
		utils.Response(w, map[string]interface{}{"status": false, "path": ""})
		return
	}
	defer f.Close()
	defer file.Close()
	utils.Response(w, map[string]interface{}{"status": true, "path": fmt.Sprintf("/upload/%s", filename)})
}
