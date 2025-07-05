package base

import (
	"github.com/stretchr/testify/require"
	"testing"
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
