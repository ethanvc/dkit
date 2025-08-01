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
			`[{"a":3}, {"a":4}]`,
			map[string]string{
				``: `a`,
			},
			`{"3":{"a":3},"4":{"a":4}}`,
		},
		{
			`[]`,
			map[string]string{},
			`[]`,
		},
	}
	for _, test := range tests {
		result, err := JsonArrayToObject([]byte(test.src), test.config)
		require.NoError(t, err, test)
		require.Equal(t, test.dst, string(result))
	}
}
