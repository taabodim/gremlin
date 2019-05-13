package gremlin

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
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

func PrepareArgs(args []interface{}, escapeFn func(string) string) []interface{} {
	for idx := range args {
		switch args[idx].(type) {
		case string:
			args[idx] = escapeFn(strings.TrimSpace(args[idx].(string)))
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
	args := PrepareArgs(gremlinQuery.Args, EscapeGremlin)
	for _, arg := range args {
		// if the argument is not a string (i.e. an int) or matches the regex string, then we're good
		if InterfaceToString(arg) != "" && argRegexP.MatchString(InterfaceToString(arg)) {
			return "", fmt.Errorf("Invalid character in your query argument: %s", InterfaceToString(arg))
		}
	}
	return fmt.Sprintf(gremlinQuery.Query, args...), nil
}

func InterfacesMatch(interface1, interface2 interface{}) bool {
	switch interface1.(type) {
	case map[string]interface{}:
		val1Map := interface1.(map[string]interface{})
		val2Map, ok := interface2.(map[string]interface{})
		if !ok || len(val1Map) != len(val2Map) {
			return false
		}
		for k, val1 := range val1Map {
			val2, ok := val2Map[k]
			if !ok {
				return false
			}
			if !InterfacesMatch(val1, val2) {
				return false
			}
		}

	case []interface{}:
		val1List := interface1.([]interface{})
		val2List, ok := interface2.([]interface{})
		if !ok || len(val1List) != len(val2List) {
			return false
		}
		for i, val1Map := range val1List {
			if !InterfacesMatch(val1Map, val2List[i]) {
				return false
			}
		}

	case []string:
		val1List := interface1.([]string)
		val2List, ok := interface2.([]string)
		if !ok || len(val1List) != len(val2List) {
			return false
		}
		for i, val1Map := range val1List {
			if !InterfacesMatch(val1Map, val2List[i]) {
				return false
			}
		}

	default:
		if interface1 == nil && interface2 != nil {
			return false
		} else if interface1 != nil && interface2 == nil {
			return false
		} else if interface1 != nil && interface2 != nil {
			if !reflect.DeepEqual(fmt.Sprintf("%v", interface1), fmt.Sprintf("%v", interface2)) {
				return false
			}
		}
	}
	return true
}
