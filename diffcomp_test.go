package dkit

import "testing"

func TestShowDiff(t *testing.T) {
	diff := NewDiffCompare()
	s1 := `{"a":"{}"}`
	s2 := `{"a":2}`
	diff.ShowDiff(s1, s2)
}
