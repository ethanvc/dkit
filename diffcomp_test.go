package dkit

import "testing"

func TestShowDiff(t *testing.T) {
	s1 := `{"a":1}`
	s2 := `{"a":2}`
	ShowDiff(s1, s2)
}
