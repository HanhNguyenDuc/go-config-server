package views

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanhnguyenduc/config-server/core"
	"github.com/hanhnguyenduc/config-server/enums"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func GetConfig(c *gin.Context) {
	var form GetConfigRequest
	if err := c.ShouldBind(&form); err != nil {
		log.Printf("[error] GetConfig | Bad request: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "Bad request",
		})
		return
	}
	log.Printf("[info] GetConfig | Request: %s", ToString(form))
	process := core.Process{
		ID:          *form.ID,
		CurrentTime: *form.CurrentTime,
		Service:     *form.Service,
	}
	isRunning, processIndex, processNum, err := core.GetResponse(process.Service, process, core.MapServiceChan[*form.Service])
	if err != nil {
		log.Printf("[error] main | %v", err)
	}
	if !isRunning {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"type":   enums.StatePending,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"type":   enums.StateRunning,
		"config": gin.H{
			"process_index": processIndex,
			"process_num":   processNum,
		},
	})
}
