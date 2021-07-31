package core

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hanhnguyenduc/config-server/conf"
	"github.com/hanhnguyenduc/config-server/enums"
	"github.com/hanhnguyenduc/config-server/redisclient"
)

var ErrStateCantBeParsed error = fmt.Errorf("redis state can't be parsed to string")

func GetCurrentState(service string) (int, error) {
	inf, err := redisclient.Get(fmt.Sprintf("%s-%s-state", enums.ConfigServerPrefix, service))
	if err != nil {
		return -1, err
	}
	log.Printf("[info] test-service current-state: %v", inf)
	res, ok := inf.(string)
	if !ok {
		return -1, ErrStateCantBeParsed
	}
	resInt, err := strconv.Atoi(res)
	if err != nil {
		return -1, err
	}
	return resInt, nil
}

func SetCurrentState(service string, state int) error {
	err := redisclient.Set(fmt.Sprintf("%s-%s-state", enums.ConfigServerPrefix, service), state, conf.RedisTimeout)
	if err != nil {
		return err
	}
	return nil
}
