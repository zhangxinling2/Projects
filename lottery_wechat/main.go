package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"lottery_wechat/config"
	"lottery_wechat/router"
)

func Init() {
	config.InitGlobalConfig()
}
func main() {
	Init()
	conf := config.GetGlobalConfig()
	fmt.Println(conf)

	log.Info("111111111111111")
	r := router.SetRouter()
	if err := r.Run(fmt.Sprintf(":%d", conf.AppConfig.Port)); err != nil {
		panic("server run fail")
	}
}
