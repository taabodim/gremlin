package gremlin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerializeVertexes(t *testing.T) {
	givens := []string{
		// test empty response
		`[]`,
		// test single vertex, single property
		`[{"@type":"g:Vertex","@value":{"id":"test-id","label":"label","properties":{"health":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"1","label":"health"}}]}}}]`,
		// test two vertexes, single property
		`[{"@type":"g:Vertex","@value":{"id":"test-id","label":"label","properties":{"health":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"1","label":"health"}}]}}}, {"@type":"g:Vertex","@value":{"id":"test-id2","label":"label","properties":{"health":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"1","label":"health"}}]}}}]`,
		// test single vertex, two properties
		`[{"@type":"g:Vertex","@value":{"id":"test-id","label":"label","properties":{"health":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"1","label":"health"}}], "health2":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"2","label":"health2"}}]}}}]`,
		// test single vertex, single property - but property has multiple values
		`[{"@type":"g:Vertex","@value":{"id":"test-id","label":"label","properties":{"health":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"1","label":"health"}}, {"@type":"g:VertexProperty","@value":{"id":{"@type":"Type","@value": 1},"value":"2","label":"health"}}]}}}]`,
	}
	expecteds := [][]Vertex{
		{},
		{MakeDummyVertex("test-id", "label", map[string]interface{}{"health": 1})},
		{MakeDummyVertex("test-id", "label", map[string]interface{}{"health": 1}), MakeDummyVertex("test-id2", "label", map[string]interface{}{"health": 1})},
		{MakeDummyVertex("test-id", "label", map[string]interface{}{"health": 1, "health2": 2})},
		{MakeDummyVertex("test-id", "label", map[string]interface{}{"health": []interface{}{1, 2}})},
	}
	for i, given := range givens {
		expected := expecteds[i]
		result, err := DeserializeVertices(given)

		if err != nil || len(result) != len(expected) {
			t.Error("given", given, "expected", expected, "result", result, "err", err)
		}
		for j, resultVertex := range result {
			expectedVertex := expected[j]

			if !VertexesMatch(resultVertex, expectedVertex) {
				t.Error("given", given, "expected", expectedVertex.Value.Properties, "result", resultVertex.Value.Properties)
			}
		}
	}
}

func TestSerializeEdges(t *testing.T) {
	givens := []string{
		// test empty response
		`[]`,
		// test single edge, single property
		`[{"@type":"g:Edge","@value":{"id":"test-id","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"inV","outV":"outV","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":"test"}}}}}]`,
		// test two edges, single property
		`[{"@type":"g:Edge","@value":{"id":"test-id","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"inV","outV":"outV","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":"test"}}}}}, {"@type":"g:Edge","@value":{"id":"test-id2","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"inV","outV":"outV","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":"test"}}}}}]`,
		// test single edge, multiple properties
		`[{"@type":"g:Edge","@value":{"id":"test-id","label":"label","inVLabel":"inVLabel","outVLabel":"outVLabel","inV":"inV","outV":"outV","properties":{"test":{"@type":"g:Property","@value":{"key":"test","value":"test"}}, "test2":{"@type":"g:Property","@value":{"key":"test2","value":1}}}}}]`,
	}
	expecteds := [][]Edge{
		{},
		{MakeDummyEdge("test-id", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test"})},
		{MakeDummyEdge("test-id", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test"}), MakeDummyEdge("test-id2", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test"})},
		{MakeDummyEdge("test-id", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test", "test2": 1})},
	}

	for i, given := range givens {
		expected := expecteds[i]
		result, err := DeserializeEdges(given)

		if err != nil || len(result) != len(expected) {
			t.Error("given", given, "expected", expected, "result", result, "err", err)
		}

		for j, resultEdge := range result {
			expectedEdge := expected[j]
			if !EdgesMatch(resultEdge, expectedEdge) {
				expectedEdgeString := fmt.Sprintf("%v", expectedEdge)
				resultEdgeString := fmt.Sprintf("%v", resultEdge)
				t.Error("given", given, "expected", expectedEdgeString, "result", resultEdgeString)
			}
		}
	}
}

func TestSerializeGenericValue(t *testing.T) {
	givens := []string{
		// test empty response
		`[]`,
		// test single gv, core return type
		`[{"@type":"g:Edge", "@value": 1}]`,
		// test 2 gv, core return type
		`[{"@type":"g:Edge", "@value": 1}, {"@type":"g:Edge2", "@value": "test"}]`,
		// test single gv, map return type
		`[{"@type":"g:Edge", "@value": {"test": "test"}}]`,
		// test single gv, nested map return type
		`[{"@type":"g:Edge", "@value": {"test": {"test": "test"}}}]`,
	}
	expecteds := [][]GenericValue{
		{},
		{MakeDummyGenericValue("g:Edge", 1)},
		{MakeDummyGenericValue("g:Edge", 1), MakeDummyGenericValue("g:Edge2", "test")},
		{MakeDummyGenericValue("g:Edge", map[string]interface{}{"test": "test"})},
		{MakeDummyGenericValue("g:Edge", map[string]interface{}{"test": map[string]interface{}{"test": "test"}})},
	}

	for i, given := range givens {
		expected := expecteds[i]
		result, err := SerializeGenericValues(given)

		if err != nil || len(result) != len(expected) {
			t.Error("given", given, "expected", expected, "result", result, "err", err)
		}

		for j, resultGenericValue := range result {
			expectedGenericValue := expected[j]
			if !GenericValuesMatch(resultGenericValue, expectedGenericValue) {
				t.Error("given", given, "expected", expectedGenericValue, "result", resultGenericValue)
			}
		}
	}
}

func TestSerializeListInterface(t *testing.T) {
	var testCases = []struct {
		given    string
		expected []interface{}
	}{
		// 1. Empty string
		{`[]`, []interface{}{}},
		// 2. Single value
		{`["test"]`, []interface{}{"test"}},
		// 3. Multiple values
		{`["test", "test2", "test3"]`, []interface{}{"test", "test2", "test3"}},
		// 4. Multiple values, mixed types
		{`["test", true, 3.12]`, []interface{}{"test", true, 3.12}},
		// 5. Multiple values, nested types
		{`[[1], {"a": 1}]`, []interface{}{[]interface{}{1}, map[string]interface{}{"a": 1}}},
	}

	for _, test := range testCases {
		result, err := SerializeListInterface(test.given)
		if err != nil {
			t.Error("given", test.given, "expected", test.expected, "result", result, "err", err)
		}
		if len(result) != len(test.expected) {
			t.Error("given", test.given, "expected", test.expected, "result", result, "err", err)
		}
		for i, r := range result {
			expectedVal := reflect.ValueOf(test.expected[i])
			resultVal := reflect.ValueOf(r)
			if !reflect.DeepEqual(expectedVal.String(), resultVal.String()) {
				t.Error("given", test.given, "expected", expectedVal.Type(), "result", resultVal, "err", err)
			}
		}
	}
}

func TestConvertToCleanVertexes(t *testing.T) {
	givens := [][]Vertex{
		{},
		{MakeDummyVertex("1", "label", map[string]interface{}{"health": 1})},
		{MakeDummyVertex("2", "label", map[string]interface{}{"health": 1}), MakeDummyVertex("2", "label", map[string]interface{}{"health": 1})},
	}
	expecteds := [][]CleanVertex{
		{},
		{CleanVertex{Id: 1, Label: "label"}},
		{CleanVertex{Id: 2, Label: "label"}, CleanVertex{Id: 2, Label: "label"}},
	}

	for i, given := range givens {
		expected := expecteds[i]
		result := ConvertToCleanVertexes(given)

		if len(result) != len(expected) {
			t.Error("given", given, "expected", expected, "result", result)
		}

		for j, resultCleanVertex := range result {
			expectedCleanVertex := expected[j]
			if expectedCleanVertex.Id != resultCleanVertex.Id || expectedCleanVertex.Label != expectedCleanVertex.Label {
				t.Error("given", given, "expected", expected, "result", result)
			}
		}
	}
}

func TestConvertToCleanEdges(t *testing.T) {
	givens := [][]Edge{
		{},
		{MakeDummyEdge("test-id", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test"})},
		{MakeDummyEdge("test-id", "label", "inVLabel", "outVLabel", "inV", "outV", map[string]interface{}{"test": "test"}), MakeDummyEdge("test-id2", "label", "inVLabel", "outVLabel", "inV2", "outV2", map[string]interface{}{"test": "test"})},
	}
	expecteds := [][]CleanEdge{
		{},
		{CleanEdge{Source: "inV", Target: "outV"}},
		{CleanEdge{Source: "inV", Target: "outV"}, CleanEdge{Source: "inV2", Target: "outV2"}},
	}

	for i, given := range givens {
		expected := expecteds[i]
		result := ConvertToCleanEdges(given)

		if len(result) != len(expected) {
			t.Error("given", given, "expected", expected, "result", result)
		}

		for j, resultCleanEdges := range result {
			expectedCleanEdges := expected[j]
			if expectedCleanEdges.Source != resultCleanEdges.Source || expectedCleanEdges.Target != expectedCleanEdges.Target {
				t.Error("given", given, "expected", expected, "result", result)
			}
		}
	}
}
