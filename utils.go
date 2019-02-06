package gremlin

import (
	"bytes"
	"fmt"
	"regexp"
)

func CharSliceToMap(chars []rune) map[rune]bool {
	charMap := make(map[rune]bool)
	for _, char := range chars {
		charMap[char] = true
	}
	return charMap
}

func InterfaceToString(i interface{}) string {
	s, _ := i.(string)
	return s
}

func CoalesceStrings(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

func EscapeArgs(args []interface{}, escapeFn func(string) string) []interface{} {
	for idx := range args {
		switch args[idx].(type) {
		case string:
			args[idx] = escapeFn(args[idx].(string))
		}
	}
	return args
}

func EscapeGremlin(value string) string {
	return escapeCharacters(value, ESCAPE_CHARS_GREMLIN)
}

func escapeCharacters(value string, escapeChars map[rune]bool) string {
	var buffer bytes.Buffer

	for _, char := range value {
		if escapeChars[char] {
			if char == pctSymbol {
				buffer.WriteRune(pctSymbol)
			} else {
				buffer.WriteRune(backslash)
			}
		}
		buffer.WriteRune(char)
	}
	return buffer.String()
}

func MakeGremlinQuery(gremlinQuery GremlinQuery, argRegexP *regexp.Regexp) (string, error) {
	args := EscapeArgs(gremlinQuery.Args, EscapeGremlin)
	for _, arg := range args {
		// if the argument is not a string (i.e. an int) or matches the regex string, then we're good
		if InterfaceToString(arg) != "" && argRegexP.MatchString(InterfaceToString(arg)) {
			return "", fmt.Errorf("Invalid character in your query argument: %s", InterfaceToString(arg))
		}
	}
	return fmt.Sprintf(gremlinQuery.Query, args...), nil
}
