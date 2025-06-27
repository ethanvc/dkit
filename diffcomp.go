package dkit

import (
	"bytes"
	"encoding/json"
	"github.com/ethanvc/dkit/base"
	"os"
	"os/exec"
	"reflect"
)

type DiffCompare struct {
	ExpandJson bool
}

func NewDiffCompare() *DiffCompare {
	return &DiffCompare{
		ExpandJson: true,
	}
}

func (com *DiffCompare) ShowDiff(content1, content2 string) error {
	content1, ext1 := com.PrepareContent(content1)
	content2, ext2 := com.PrepareContent(content2)
	f1, err := os.CreateTemp("", "dkit_temp_*."+ext1)
	if err != nil {
		return err
	}
	f1.WriteString(content1)
	f1.Close()
	f2, err := os.CreateTemp("", "dkit_temp_*."+ext2)
	if err != nil {
		return err
	}
	f2.WriteString(content2)
	f2.Close()
	cmd := exec.Command("code", "--diff", f1.Name(), f2.Name())
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (com *DiffCompare) PrepareContent(content string) (string, string) {
	if json.Valid([]byte(content)) {
		return com.prepareJson(content)
	}
	return content, "txt"
}

func (com *DiffCompare) prepareJson(content string) (string, string) {
	const ext = "json"
	if !com.ExpandJson {
		buf := bytes.NewBuffer(nil)
		json.Indent(buf, []byte(content), "", "    ")
		return buf.String(), ext
	}
	var data any
	json.Unmarshal([]byte(content), &data)
	walker := base.NewObjWalker()
	walker.Walk(data, func(obj, key, val reflect.Value) (base.VisitResult, reflect.Value) {
		keyAny := key.Interface()
		valAny := val.Interface()
		_ = keyAny
		_ = valAny
		strVal, ok := reflectGetStr(val)
		if !ok {
			return base.VisitResultContinue, val
		}
		if !json.Valid([]byte(strVal)) {
			return base.VisitResultContinue, val
		}
		var anyVal any
		json.Unmarshal([]byte(strVal), &anyVal)
		return base.VisitResultContinue, reflect.ValueOf(anyVal)
	})
	newContent, _ := json.MarshalIndent(data, "", "    ")
	return string(newContent), ext
}

func reflectGetStr(val reflect.Value) (string, bool) {
	for {
		kind := val.Kind()
		if kind == reflect.String {
			return val.String(), true
		}
		if kind == reflect.Interface {
			val = val.Elem()
			continue
		}
		return "", false
	}
}
