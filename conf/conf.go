package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/houzhongjian/jiancheng/base"
)

var Conf *base.Conf

//ConfInit 初始化配置文件.
func ConfInit() {
	b, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	conf := &base.Conf{}
	if err := json.Unmarshal(b, &conf); err != nil {
		log.Printf("%+v\n", err)
		return
	}

	Conf = conf
}
