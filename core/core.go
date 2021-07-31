package core

import (
	"fmt"
	"log"

	"github.com/hanhnguyenduc/config-server/conf"
	"github.com/hanhnguyenduc/config-server/enums"
	"github.com/hanhnguyenduc/config-server/redisclient"
)

var Services []string = []string{
	"test-service",
	"smartivr-scan-campaign",
	"smartivr-scan-call",
}

var MapServiceChan map[string]chan int

func Setup() {
	// state, listProcess
	MapServiceChan = make(map[string]chan int)
	for _, service := range Services {
		redisclient.Set(
			fmt.Sprintf("%s-%s-state", enums.ConfigServerPrefix, service),
			enums.StateRunning,
			conf.RedisTimeout,
		)
		MapServiceChan[service] = make(chan int, 1)

		log.Printf("[info] core.Setup | chan: %p", MapServiceChan[service])
		go ServiceManage(service, MapServiceChan[service])
		// // Put card to mapServiceChan
		MapServiceChan[service] <- 1
	}
	log.Printf("[info] core.Setup | done")
}
