package stringtools_test

import (
	"testing"

	. "github.com/halliday/go-tools/stringtools"
)

func TestCamelToSnake(t *testing.T) {

	var camels = [][]string{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"aB", "a_b"},
		{"AB", "ab"},
		{"aBC", "a_bc"},
		{"ABC", "abc"},
		{"fooBAR", "foo_bar"},
		{"FooBarBatz", "foo_bar_batz"},
	}

	for _, c := range camels {
		snake := CamelToSnake(c[0])
		if snake != c[1] {
			t.Errorf("CamelToSnake(%q) == %q, want %q", c[0], snake, c[1])
		}
	}
}
