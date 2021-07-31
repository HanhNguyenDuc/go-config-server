package core

import (
	"log"

	"github.com/hanhnguyenduc/config-server/enums"
)

// Return isRuning: bool, processID: int, processNum: int
func GetResponse(service string, requestedProcess Process, c chan int) (bool, int, int, error) {
	<-c
	processIndex := 0
	processNum := 0
	processes, err := GetProcessesFromRedis(service)
	if err != nil {
		c <- 1
		return false, -1, -1, err
	}
	log.Printf("[info] GetResponse | service: %s | pre-processes: %v", service, processes)
	isAvailable := false
	for id, process := range processes {
		if requestedProcess.ID == process.ID {
			processes[id].CurrentTime = requestedProcess.CurrentTime
			processes[id].IsPending = false
			processIndex = id
			processNum = len(processes)
			isAvailable = true
			break
		}
	}
	log.Printf("[info] GetResponse | new_proccesses_list: %v", processes)
	// if process is not in list, append it to list
	if !isAvailable {
		requestedProcess.IsPending = true
		processes = append(processes, requestedProcess)
		log.Printf("[info] GetResponse | service: %s | post-processes: %v", service, processes)
		processes, err = SetAllProcessToPending(service, processes)
		if err != nil {
			c <- 1
			return false, -1, -1, err
		}
		err := SetCurrentState(service, enums.StatePending)

		if err != nil {
			c <- 1
			return false, -1, -1, err
		}
		err = SetProcessToRedis(service, processes)
		if err != nil {
			c <- 1
			return false, -1, -1, err
		}
		c <- 1
		return false, -1, -1, nil
	} else {
		// if process in list, change its status in redis
		err = SetProcessToRedis(service, processes)
		if err != nil {
			c <- 1
			return false, -1, -1, nil
		}
		someonePending := false
		for _, process := range processes {
			if process.IsPending {
				someonePending = true
			}
		}
		if someonePending {
			c <- 1
			return false, -1, -1, nil
		}
		err := SetCurrentState(service, enums.StateRunning)
		if err != nil {
			c <- 1
			return false, -1, -1, err
		}
		c <- 1
		return true, processIndex, processNum, nil
	}
}
