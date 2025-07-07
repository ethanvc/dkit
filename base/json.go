package base

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func ExpandJson(src []byte) []byte {
	w := NewJsonWalker()
	dst := src
	var err error
	w.Walk(src, func(p JsonWalkPath, val gjson.Result) JsonVisitResult {
		if val.Type != gjson.String {
			return JsonVisitResultContinue
		}
		if !isArrayOrObject(val.String()) {
			return JsonVisitResultContinue
		}
		expandedChild := ExpandJson([]byte(val.String()))
		childPath := p.GetFullPath()
		dst, err = sjson.SetRawBytes(dst, childPath, expandedChild)
		if err != nil {
			panic("why have error")
		}
		return JsonVisitResultContinue
	})
	return dst
}

// to gjson, "[Linked] 6365" string is recgnized as an array
func isArrayOrObject(val string) bool {
	if !json.Valid([]byte(val)) {
		return false
	}
	result := gjson.Parse(val)
	return result.IsObject() || result.IsArray()
}
