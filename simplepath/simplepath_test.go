package simplepath

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimplePath_Parse(t *testing.T) {
	var p SimplePath
	var err error
	p, err = Parse("a.b.c")
	require.NoError(t, err)
	require.Equal(t, `a.b.c`, p.String())
}
