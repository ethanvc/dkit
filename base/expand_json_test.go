package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpandJson(t *testing.T) {
	src := `{"a":"{}"}`
	dst, err := ExpandJson([]byte(src))
	require.NoError(t, err)
	require.Equal(t, `{"a":{}}`, string(dst))
}

func TestExpandJson2(t *testing.T) {
	src := `["{}", "{}"]`
	dst, err := ExpandJson([]byte(src))
	require.NoError(t, err)
	require.Equal(t, `[{},{}]`, string(dst))
}

func TestExpandJson3(t *testing.T) {
	src := `"[\"{}\", \"{}\"]"`
	dst, err := ExpandJson([]byte(src))
	require.NoError(t, err)
	require.Equal(t, `[{},{}]`, string(dst))
}
