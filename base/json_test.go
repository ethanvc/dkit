package base

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpandJson(t *testing.T) {
	src := `{"a":"{}"}`
	dst := ExpandJson([]byte(src))
	require.Equal(t, ``, string(dst))
}

func TestExpandJson2(t *testing.T) {
	src := `["{}", "{}"]`
	dst := ExpandJson([]byte(src))
	require.Equal(t, `[ { }, { } ]`, string(dst))
}

func TestExpandJsonFromFile(t *testing.T) {
	content, err := os.ReadFile("1.json")
	require.NoError(t, err)
	newContent := ExpandJson(content)
	valid := json.Valid(newContent)
	require.True(t, valid)
}
