package base

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type keyType string

const (
	keyTypeSuffix keyType = "suffix"
	KeyTypeFull   keyType = "full"
)

func JsonArrayToObject(data []byte, config map[string]string) []byte {
	newData := data
	w := NewJsonWalker()
	w.Walk(data, func(p JsonWalkPath, val gjson.Result) JsonVisitResult {
		if !val.IsArray() {
			return JsonVisitResultContinue
		}
		valKeyPath := config[p.GetLastNode()]
		if valKeyPath == "" {
			return JsonVisitResultContinue
		}
		return JsonVisitResultContinue
	})
	return newData
}

func convertArrayToObject(data []byte, sortKey string) []byte {
	result := gjson.ParseBytes(data)
	newData := []byte("{}")
	var err error
	for _, item := range result.Array() {
		val := item.Get(sortKey)
		if val.String() == "" {
			return data
		}
		newData, err = sjson.SetRawBytes(newData, val.String(), []byte(item.Raw))
		if err != nil {
			fmt.Println(err)
		}
	}
	return newData
}
