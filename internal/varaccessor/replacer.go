package varaccessor

import (
	"regexp"
	"strings"
)

// variable replacer with regex: {{ var.name.and.subname }}
var varReplacer = regexp.MustCompile(`\{\{\s*([a-zA-Z0-9_]+)\s*\}\}`)

// ReplaceVars replaces variables in string
func ReplaceVars(s string, vars map[string]any) string {
	return ReplaceVarsWithRe(s, vars, varReplacer)
}

// ReplaceVarsWithRe replaces variables in string with custom regexp
func ReplaceVarsWithRe(s string, vars map[string]any, re *regexp.Regexp) string {
	varAccessor := New(vars)
	return re.ReplaceAllStringFunc(s, func(s string) string {
		submatch := re.FindStringSubmatch(s)
		if len(submatch) > 1 {
			return varAccessor.String(submatch[1], s)
		}
		return varAccessor.String(strings.TrimSpace(s[2:len(s)-2]), s)
	})
}
