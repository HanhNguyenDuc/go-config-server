package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanhnguyenduc/config-server/views"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/getConfig", views.GetConfig)
	r.GET("/", views.Index)
	return r
}
