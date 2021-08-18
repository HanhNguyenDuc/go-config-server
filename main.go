package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanhnguyenduc/config-server/core"
	"github.com/hanhnguyenduc/config-server/redisclient"
	"github.com/hanhnguyenduc/config-server/routes"
	"github.com/hanhnguyenduc/config-server/setting"
)

func main() {
	isDeploy := flag.Bool(
		"deploy",
		false,
		"to know this app is deploy or not, if this app is deploying, load config from envs instead",
	)
	flag.Parse()
	setting.Setup(isDeploy)
	flag.Parse()
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
