package base

import (
	"encoding/json"
	"github.com/ethanvc/dkit/simplepath"
)

func JsonArrayToObject(data []byte, config map[string]string) ([]byte, error) {
	var dataObj any
	err := Unmarshal(data, &dataObj)
	if err != nil {
		return nil, err
	}
	dataObj, err = JsonAnyArrayToObject(dataObj, config)
	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(dataObj)
	return data, err
}

func JsonAnyArrayToObject(data any, config map[string]string) (any, error) {
	return jsonAnyArrayToObject(data, nil, config)
}

func jsonAnyArrayToObject(data any, p simplepath.SimplePath, config map[string]string) (any, error) {
	realData, ok := data.([]any)
	if !ok {
		return data, nil
	}
	return nil, nil
}

func findMatchedConfig(config map[string]string, currentP simplepath.SimplePath) (string, bool) {
	for p, key := range config {

	}
}
