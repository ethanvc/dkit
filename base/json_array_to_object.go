package base

import (
	"encoding/json"
	"fmt"
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
	switch realData := data.(type) {
	case map[string]any:
		for k, v := range realData {
			v, err := jsonAnyArrayToObject(v, p.AppendKey(k), config)
			if err != nil {
				return nil, err
			}
			realData[k] = v
		}
		return data, nil
	case []any:
		for i, v := range realData {
			v, err := jsonAnyArrayToObject(v, p.AppendIndex(i), config)
			if err != nil {
				return nil, err
			}
			realData[i] = v
		}
		return convertArrayToObject(realData, p, config)
	default:
		return data, nil
	}
}

func convertArrayToObject(arrayData []any, p simplepath.SimplePath, config map[string]string) (any, error) {
	key, ok := findMatchedConfig(config, p)
	if !ok {
		return arrayData, nil
	}
	result := make(map[string]any, len(arrayData))
	keyPath, err := simplepath.Parse(key)
	if err != nil {
		return nil, err
	}
	for _, data := range arrayData {
		keyVal, _ := keyPath.Get(data)
		keyValStr := fmt.Sprintf("%v", keyVal)
		result[keyValStr] = data
	}
	return result, nil
}

func findMatchedConfig(config map[string]string, currentP simplepath.SimplePath) (string, bool) {
	for p, key := range config {
		pattern, err := simplepath.Parse(p)
		if err != nil {
			continue
		}
		if currentP.Match(pattern) {
			return key, true
		}
	}
	return "", false
}
