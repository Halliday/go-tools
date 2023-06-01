package stringtools

import (
	"regexp"
	"strings"
)

var camelRegexp = regexp.MustCompile(`[A-Z]+`)

// CamelToSnake replaces camel case with snake case.
// For example, "fooBar" becomes "foo_bar".
// More examples:
//
//	"FooBar" -> "foo_bar"
//	"FooBarBatz" -> "foo_bar_batz"
//	"fooBAR" -> "foo_bar"
func CamelToSnake(str string) string {
	return strings.TrimLeft(camelRegexp.ReplaceAllStringFunc(str, camelCaseString), "_")
}

func camelCaseString(s string) string {
	return "_" + strings.ToLower(s)
}
