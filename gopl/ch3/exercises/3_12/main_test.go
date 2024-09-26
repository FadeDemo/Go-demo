package main

import (
	"testing"
)

func TestIsIsomerization(t *testing.T) {
	var examples = []struct {
		s1, s2 string
		want   bool
	}{
		{"abc", "bca", true},
		{"a你好", "你a好", true},
		{"abc", "你好a", false},
	}
	for _, ex := range examples {
		if got := IsIsomerization(ex.s1, ex.s2); got != ex.want {
			t.Errorf("%s and %s got %v, want %v", ex.s1, ex.s2, got, ex.want)
		}
	}
}
