package dkit

import (
	"bytes"
	"encoding/json"
	"github.com/ethanvc/dkit/base"
	"os"
	"os/exec"
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
	f1, err := os.CreateTemp("", "dkit_benchmark_*."+ext1)
	if err != nil {
		return err
	}
	f1.WriteString(content1)
	f1.Close()
	f2, err := os.CreateTemp("", "dkit_target_*."+ext2)
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
	expandContent := base.ExpandJson([]byte(content))
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, expandContent, "", "    ")
	if err != nil {
		panic(err)
	}
	return buf.String(), ext
}
