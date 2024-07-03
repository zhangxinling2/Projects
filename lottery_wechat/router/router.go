package router

import (
	"github.com/gin-gonic/gin"
	"lottery_wechat/api"
)

func SetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	group := r.Group("/lottery_wechat")
	group.GET("/hello", api.Hello)
	return r
}
