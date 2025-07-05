package base

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
)

type JsonWalker struct {
}

func NewJsonWalker() *JsonWalker {
	return &JsonWalker{}
}

type JsonWalkPath []any

func (p JsonWalkPath) AppendKey(key string) JsonWalkPath {
	return append(p, key)
}

func (p JsonWalkPath) AppendIndex(index int) JsonWalkPath {
	return append(p, index)
}

func (p JsonWalkPath) GetFullPath() string {
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

type JsonVisitFunc func(p JsonWalkPath, val gjson.Result) JsonVisitResult

type JsonVisitResult int

const (
	JsonVisitResultContinue JsonVisitResult = iota
	JsonVisitResultStop
	JsonVisitResultSkipCurrentValue
)

func (w *JsonWalker) Walk(content []byte, cb JsonVisitFunc) {
	obj := gjson.ParseBytes(content)
	w.walk(obj, nil, cb)
}

func (w *JsonWalker) walk(obj gjson.Result, p JsonWalkPath, fn JsonVisitFunc) bool {
	obj.ForEach(func(key, val gjson.Result) bool {
		switch key.Type {
		case gjson.String:
			childP := p.AppendKey(key.String())
			fn(childP, val)
		case gjson.Number:
			childP := p.AppendIndex(int(key.Int()))
			fn(childP, val)
		default:
			fmt.Println("unsupported key type")
		}
		return true
	})
	return true
}
