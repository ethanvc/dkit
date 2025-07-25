package dgit

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_runCommand(t *testing.T) {
	c := context.Background()
	buf, _, err := runCommand(c, "git", "rev-parse", "--is-inside-work-tree")
	require.NoError(t, err)
	_ = buf
}
