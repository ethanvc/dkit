package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJsonArrayToObject(t *testing.T) {
	type args struct {
		src    string
		config map[string]string
		dst    string
	}

	tests := []args{
		{
			"[]",
			map[string]string{},
			"[]",
		},
	}
	for _, test := range tests {
		result, err := JsonArrayToObject([]byte(test.src), test.config)
		require.NoError(t, err, test)
		require.Equal(t, test.dst, string(result))
	}
}
