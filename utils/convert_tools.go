package utils

import (
	"encoding/json"
	"fmt"
)

var ErrCantParseToString error = fmt.Errorf("cant parse interface{} to string")

func ParseInterfaceToList(inf interface{}, des interface{}) error {

	JSON, ok := inf.(string)
	if !ok {
		return ErrCantParseToString
	}
	err := json.Unmarshal([]byte(JSON), des)
	if err != nil {
		return err
	}
	return nil
}

func ParseListToInterface(list interface{}) (interface{}, error) {
	JSON, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return string(JSON), nil
}
