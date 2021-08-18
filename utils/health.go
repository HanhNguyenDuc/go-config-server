package utils

import "io/ioutil"

func ConfirmHealthy() error {
	healthFileContent := "2000"
	err := ioutil.WriteFile("health/healthy", []byte(healthFileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ConfirmUnhealthy() error {
	unhealthFileContent := "2000"
	err := ioutil.WriteFile("health/unhealthy", []byte(unhealthFileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}
