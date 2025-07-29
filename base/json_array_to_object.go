package base

import "encoding/json"

func JsonArrayToObject(data []byte, config map[string]string) ([]byte, error) {
	var dataObj any
	err := Unmarshal(data, &dataObj)
	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(dataObj)
	return data, err
}

type jsonArrayToObjectConfig struct {
	config map[string][]string
}

func newConfig(config map[string]string) jsonArrayToObjectConfig {

}
