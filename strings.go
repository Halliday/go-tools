package tools

import (
	"regexp"
	"strings"
)

var camelRegexp = regexp.MustCompile(`.?[A-Z]`)

func CamelToSnake(camel string) string {
	var snake = camelRegexp.ReplaceAllStringFunc(camel, func(s string) string {
		if len(s) == 1 {
			return strings.ToLower(s)
		}
		return s[:1] + "_" + strings.ToLower(s[1:])
	})
	return snake
}
