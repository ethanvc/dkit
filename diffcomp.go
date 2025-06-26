package dkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func ShowDiff(content1, content2 string) {
	content1, ext1 := beautifyContent(content1)
	content2, ext2 := beautifyContent(content2)
	f1, err := os.CreateTemp("", "dkit_temp_*."+ext1)
	if err != nil {
		panic(err)
	}
	f1.WriteString(content1)
	f1.Close()
	f2, err := os.CreateTemp("", "dkit_temp_*."+ext2)
	if err != nil {
		panic(err)
	}
	f2.WriteString(content2)
	f2.Close()
	cmd := exec.Command("code", "--diff", f1.Name(), f2.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Printf("start code command failed: %s\n", err.Error())
	}
}

func beautifyContent(content string) (string, string) {
	if json.Valid([]byte(content)) {
		buf := bytes.NewBuffer(nil)
		json.Indent(buf, []byte(content), "", "    ")
		return buf.String(), "json"
	}
	return content, "txt"
}
