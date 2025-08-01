package base

import (
	"encoding/json"
	"net/url"
)

func ExpandJson(data []byte) ([]byte, error) {
	var jsonObj any
	err := Unmarshal(data, &jsonObj)
	if err != nil {
		return nil, err
	}
	jsonObj, err = ExpandJsonAny(jsonObj)
	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(jsonObj)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ExpandJsonAny(data any) (any, error) {
	switch realData := data.(type) {
	case map[string]any:
		for k, v := range realData {
			newV, err := ExpandJsonAny(v)
			if err != nil {
				return nil, err
			}
			realData[k] = newV
		}
	case []any:
		for i, item := range realData {
			newData, err := ExpandJsonAny(item)
			if err != nil {
				return nil, err
			}
			realData[i] = newData
		}
	case string:
		if json.Valid([]byte(realData)) {
			var newData any
			err := Unmarshal([]byte(realData), &newData)
			if err != nil {
				return nil, err
			}
			return ExpandJsonAny(newData)
		}
		if result, ok := expandUrl(realData); ok {
			return result, nil
		}
	}
	return data, nil
}

func expandUrl(data string) (any, bool) {
	u, err := url.Parse(data)
	if err != nil {
		return nil, false
	}
	mapData := make(map[string]any)
	queries := u.Query()
	u.RawQuery = ""
	mapData["url"] = u.String()
	mapData["queries"] = queries
	return mapData, true
}
