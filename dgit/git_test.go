package dgit

import (
	"context"
	"testing"
)

func Test_runCommand(t *testing.T) {
	c := context.Background()
	runCommand(c, "git", "rev-parse", "--is-inside-work-tree")
}
