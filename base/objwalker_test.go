package base

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewObjWalker(t *testing.T) {
	walker := NewObjWalker()
	{
		obj := map[string]any{
			"hello": "3",
		}
		walker.Walk(obj, func(obj, key, val reflect.Value) (VisitResult, reflect.Value) {
			return VisitResultContinue, reflect.ValueOf("33")
		})
		require.Equal(t, "33", obj["hello"].(string))
	}
	{
		obj := []string{"1", "2"}
		walker.Walk(obj, func(obj, key, val reflect.Value) (VisitResult, reflect.Value) {
			return VisitResultContinue, reflect.ValueOf("33")
		})
		require.Equal(t, "33", obj[0])
	}
}
