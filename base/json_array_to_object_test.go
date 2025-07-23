package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_convertArrayToObject(t *testing.T) {
	data := []byte(`[{"index":1}, {"index":2}]`)
	newData := convertArrayToObject(data, "index")
	require.Equal(t, `{"1":{"index":1},"2":{"index":2}}`, string(newData))
}
