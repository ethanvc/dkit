package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func Unmarshal(data []byte, v any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

type JsonPath []any

func (p JsonPath) AppendKey(key string) JsonPath {
	return append(p, key)
}

func (p JsonPath) AppendIndex(index int) JsonPath {
	return append(p, index)
}

func (p JsonPath) GetLastNode() string {
	if len(p) == 0 {
		return ""
	}
	v := p[len(p)-1]
	switch realV := v.(type) {
	case string:
		return realV
	case int:
		return strconv.Itoa(realV)
	}
	panic("unreachable")
}

func (p JsonPath) GetFullPath() string {
	buf := bytes.NewBuffer(nil)
	for _, v := range p {
		switch realV := v.(type) {
		case string:
			if buf.Len() == 0 {
				buf.WriteString(realV)
				continue
			}
			buf.WriteString(fmt.Sprintf(".%s", realV))
		case int:
			if buf.Len() == 0 {
				buf.WriteString(fmt.Sprintf("%d", realV))
				continue
			}
			buf.WriteString(fmt.Sprintf(".%d", realV))
		default:
			panic("unsupported type")
		}
	}
	return buf.String()
}
