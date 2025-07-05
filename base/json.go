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
		switch val.Type {
		case gjson.String:
			if json.Valid([]byte(val.String())) {
				expandedChild := ExpandJson([]byte(val.String()))
				childPath := p.GetFullPath()
				dst, err = sjson.SetRawBytes(dst, childPath, expandedChild)
				if err != nil {
					panic("why have error")
				}
				return JsonVisitResultContinue
			}
			return JsonVisitResultStop
		}
		return JsonVisitResultContinue
	})
	return dst
}
