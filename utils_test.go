package gremlin

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

func TestCharSliceToMap(t *testing.T) {
	given := []rune{
		singleQuote,
		backslash,
		pctSymbol,
		doubleQuote,
	}
	expectedMap := make(map[rune]bool)
	expectedMap[singleQuote] = true
	expectedMap[backslash] = true
	expectedMap[pctSymbol] = true
	expectedMap[doubleQuote] = true
	expected, _ := json.Marshal(expectedMap)

	resultMap := CharSliceToMap(given)
	result, _ := json.Marshal(resultMap)
	if string(result) != string(expected) {
		t.Error("given", given, "expected", expected, "result", result)
	}
}

func TestInterfaceToString(t *testing.T) {
	tests := [][]interface{}{
		{"", ""},
		{"test", "test"},
		{1, ""},
		{true, ""},
	}

	for _, test := range tests {
		given := test[0]
		expected := test[1].(string)
		result := InterfaceToString(given)
		if result != expected {
			t.Error("given", given, "expected", expected, "result", result)
		}
	}
}

func TestCoalesceStrings(t *testing.T) {
	tests := [][][]string{
		{{"first"}, {"first"}},
		{{"", "first"}, {"first"}},
		{{"", "first", "", "second"}, {"first"}},
	}
	for _, test := range tests {
		given := test[0]
		expected := test[1][0]
		result := CoalesceStrings(given...)
		if result != expected {
			t.Error("given", given, "expected", expected, "result", result)
		}
	}
}

func TestEscapeCharacters(t *testing.T) {
	tests := [][]string{
		{"this is a test", "this is a test"},
		{`this is a %`, `this is a %%`},
		{"", ""},
		{`' \ % "`, `\' \\ %% \"`},
	}
	for _, test := range tests {
		given := test[0]
		expected := test[1]
		result := escapeCharacters(given, ESCAPE_CHARS_GREMLIN)
		if result != expected {
			t.Error("given", given, "expected", expected, "result", result)
		}
	}
}

func TestEscapeGremlin(t *testing.T) {
	tests := [][]string{
		{"this is a test", "this is a test"},
		{`this is a %`, `this is a %%`},
		{"", ""},
		{`' \ % "`, `\' \\ %% \"`},
	}
	for _, test := range tests {
		given := test[0]
		expected := test[1]
		result := EscapeGremlin(given)
		if result != expected {
			t.Error("given", given, "expected", expected, "result", result)
		}
	}
}

func TestPrepareArgs(t *testing.T) {
	tests := [][][]interface{}{
		{{"blah"}, {"blah"}},
		{{"blah", 1}, {"blah", 1}},
		{{"blah", 1, `escape '`}, {"blah", 1, `escape \'`}},
		{{`	`, ` this`, `this `, ` this `}, {``, `this`, `this`, `this`}},
	}
	for _, test := range tests {
		given := test[0]
		expected := fmt.Sprintf("%v", test[1])
		result := fmt.Sprintf("%v", PrepareArgs(given, EscapeGremlin))
		if result != expected {
			t.Error("given", given, "expected", expected, "result", result)
		}
	}
}

func TestMakeGremlinQuery(t *testing.T) {
	vertexId := "vertexId"
	label := "label"
	currentTime := "test"
	testCases := []struct {
		given    GremlinQuery
		expected string
	}{
		// 1. Very basic test
		{
			GremlinQuery{
				Query: "g.V('%s')",
				Args:  []interface{}{"test"},
			},
			"g.V('test')",
		},
		// 2. Basic test using one of our upsert queries
		{
			GremlinQuery{
				Query: `g.V("%s").fold().coalesce(unfold().coalesce(has("hidden").drop(), g.V("%s")),g.addV("%s").property(id, "%s").property(single, "create_date", "%s")).property(single, "last_updated_date", "%s").property(single, "%v", "%v")`,
				Args:  []interface{}{vertexId, vertexId, label, vertexId, currentTime, currentTime, "prop_1", "prop_2"},
			},
			fmt.Sprintf(`g.V("%s").fold().coalesce(unfold().coalesce(has("hidden").drop(), g.V("%s")),g.addV("%s").property(id, "%s").property(single, "create_date", "%s")).property(single, "last_updated_date", "%s").property(single, "%v", "%v")`, vertexId, vertexId, label, vertexId, currentTime, currentTime, "prop_1", "prop_2"),
		},
		// 3. Test using %
		{
			GremlinQuery{
				Query: `g.V("%s").property(single, "%s", "%s")`,
				Args:  []interface{}{"test%", "te%st", "\"%this"},
			},
			`g.V("test%%").property(single, "te%%st", "\"%%this")`,
		},
		// 4. Test using the rest of the punctuation
		{
			GremlinQuery{
				Query: `g.V("%s")`,
				Args:  []interface{}{"test/-?!*()&_=,#'><+@;.:$\\"},
			},
			`g.V("test/-?!*()&_=,#\'><+@;.:$\\")`,
		},
		// 5. Test using a big prop from our system
		{
			GremlinQuery{
				Query: `g.V("%s").property(single, "%s", "%s")`,
				Args:  []interface{}{"test", "query", `SELECT prop1 as proP1, prop2, prop3 as proP3, prop4 as prop$ FROM dummy.dummy WHERE some_id <> 111111 AND other LIKE \'%@test.test\' AND prop9 = 1 AND (prop_10 > 0 OR prop_11 >= NOW() - INTERVAL 1 WEEK)`},
			},
			`g.V("test").property(single, "query", "SELECT prop1 as proP1, prop2, prop3 as proP3, prop4 as prop$ FROM dummy.dummy WHERE some_id <> 111111 AND other LIKE \\\'%%@test.test\\\' AND prop9 = 1 AND (prop_10 > 0 OR prop_11 >= NOW() - INTERVAL 1 WEEK)")`,
		},
	}
	argRegexp, _ := regexp.Compile(ARG_REGEX)

	for _, test := range testCases {
		result, err := MakeGremlinQuery(test.given, argRegexp)

		if err != nil || result != test.expected {
			t.Error("given", test.given, "expected", test.expected, "result", result, "err", err)
		}
	}
}

func TestMakeGremlinQueryError(t *testing.T) {
	// Testing characters we don't allow
	argRegexp, _ := regexp.Compile(ARG_REGEX)

	testCases := []struct {
		given GremlinQuery
	}{
		// 1. Cc - Basic control character
		{
			GremlinQuery{
				Query: "g.V('%s')",
				Args:  []interface{}{fmt.Sprintf("test] %s", string('\x00'))},
			},
		},
		// 2. Cf - formatting indicator
		{
			GremlinQuery{
				Query: "g.V('%s')",
				Args:  []interface{}{fmt.Sprintf("test] %s", string('\u00AD'))},
			},
		},
		// 3. Co - private code point
		{
			GremlinQuery{
				Query: "g.V('%s')",
				Args:  []interface{}{fmt.Sprintf("test] %s", string('\uE000'))},
			},
		},
		// 4. Cn - Unassigned code point that can break Neptune
		{
			GremlinQuery{
				Query: "g.V('%s')",
				Args:  []interface{}{fmt.Sprintf("test] %s", string('\U0010e4c3'))},
			},
		},
	}

	for _, test := range testCases {
		result, err := MakeGremlinQuery(test.given, argRegexp)
		if err == nil {
			t.Error("given", test.given, "result", result, "err", err)
		}
	}
}
