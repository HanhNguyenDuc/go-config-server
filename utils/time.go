package utils

import "time"

// GetCurrentTimeByMilisecond ...
func GetCurrentTimeByMilisecond() int64 {
	now := time.Now()
	unixNano := now.UnixNano()
	milisec := unixNano / 1000000
	return milisec
}
