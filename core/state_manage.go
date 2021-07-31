package core

import (
	"fmt"
	"log"
	"time"

	"github.com/hanhnguyenduc/config-server/conf"
	"github.com/hanhnguyenduc/config-server/enums"
	"github.com/hanhnguyenduc/config-server/redisclient"
	"github.com/hanhnguyenduc/config-server/utils"
)

type Process struct {
	ID          string `json:"id"`
	CurrentTime int64  `json:"current_time"`
	Service     string `json:"service"`
	IsPending   bool
}

func GetProcessesFromRedis(service string) ([]Process, error) {
	serviceProcessesInf, err := redisclient.Get(fmt.Sprintf("%s-%s-processes", enums.ConfigServerPrefix, service))
	if err != nil && err.Error() != "redis: nil" {

		return nil, err
	}
	if serviceProcessesInf == nil {
		return make([]Process, 0), nil
	}
	serviceProcesses := make([]Process, 0)
	err = utils.ParseInterfaceToList(serviceProcessesInf, &serviceProcesses)
	if err != nil {
		return nil, err
	}
	return serviceProcesses, nil
}

func SetProcessToRedis(service string, processes []Process) error {
	processJSON, err := utils.ParseListToInterface(processes)
	if err != nil {
		return err
	}
	err = redisclient.Set(
		fmt.Sprintf("%s-%s-processes", enums.ConfigServerPrefix, service),
		processJSON,
		conf.RedisTimeout,
	)
	if err != nil {
		return err
	}
	return nil
}

func SetAllProcessToPending(service string, processes []Process) ([]Process, error) {
	// processes, err := GetProcessesFromRedis(service)
	var err error
	for id := range processes {
		processes[id].IsPending = true
	}
	err = SetProcessToRedis(service, processes)
	if err != nil {
		return nil, err
	}
	return processes, nil
}

func RemoveProcessByIds(processes []Process, ids []int) []Process {
	res := make([]Process, 0)
	for i := 0; i < len(processes); i++ {
		flag := true
		for _, id := range ids {
			if i == id {
				flag = false
			}
		}
		if flag {
			res = append(res, processes[i])
		}
	}
	return res
}

// ServiceManage ... Detect dead process
func ServiceManage(service string, c chan int) {
	// Detect dead process
	// Change state from pending to running
	for {
		time.Sleep(5 * time.Second)
		<-c
		// Get list of processes from redis
		processes, err := GetProcessesFromRedis(service)
		if err != nil {
			log.Printf("[error] ServiceManage | service: %s | %v", service, err)
			c <- 1
			continue
		}
		// Get current time
		curTime := utils.GetCurrentTimeByMilisecond()
		removeIds := make([]int, 0)
		flag := false
		for id, process := range processes {
			// Check if process has reach timeout or not
			if curTime > process.CurrentTime+conf.MaxTimeToResponse {
				// Process is died
				err := SetCurrentState(service, enums.StatePending)
				if err != nil {
					log.Printf("[error] ServiceManage | service: %s | %v", service, err)
					flag = true
					break
				}
				log.Printf("[info] ServiceManage | service: %v | process %s has been dead", service, process.ID)
				removeIds = append(removeIds, id)
			}
		}
		if flag {
			c <- 1
			continue
		}
		err = SetCurrentState(service, enums.StatePending)
		if err != nil {
			log.Printf("[error] ServiceManage | service: %s | %v", service, err)
			c <- 1
			continue
		}
		if len(removeIds) > 0 {
			// Delete process from list
			processes = RemoveProcessByIds(processes, removeIds)

			processes, err = SetAllProcessToPending(service, processes)
			if err != nil {
				log.Printf("[error] ServiceManage | service: %s | %v", service, err)
				c <- 1
				continue
			}
			log.Printf("[info] ServiceManage | removedProcesses: %v", processes)
			// Write service processes to redis
			err = SetProcessToRedis(service, processes)
			if err != nil {
				log.Printf("[error] ServiceManage | service: %s | %v", service, err)
				c <- 1
				continue
			}
		}
		c <- 1
	}
}
