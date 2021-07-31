package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanhnguyenduc/config-server/core"
	"github.com/hanhnguyenduc/config-server/redisclient"
	"github.com/hanhnguyenduc/config-server/routes"
	"github.com/hanhnguyenduc/config-server/setting"
)

func main() {
	redisclient.Setup()
	core.Setup()
	// Define server
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routes.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
