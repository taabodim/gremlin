package gremlin

import (
	"testing"
)

func TestGenericValuesEquals(t *testing.T) {
	var testCases = []struct {
		given1   GenericValue
		given2   GenericValue
		expected bool
	}{
		{
			given1:   GenericValue{},
			given2:   GenericValue{},
			expected: true,
		},
		{
			given1:   GenericValue{"test", "test"},
			given2:   GenericValue{"test", "test"},
			expected: true,
		},
		{
			given1:   GenericValue{"test", 1},
			given2:   GenericValue{"test", 1},
			expected: true,
		},
		{
			given1:   GenericValue{"test", map[string]interface{}{"test": "test"}},
			given2:   GenericValue{"test", map[string]interface{}{"test": "test"}},
			expected: true,
		},
		{
			given1:   GenericValue{"test", "test"},
			given2:   GenericValue{"test", "test1"},
			expected: false,
		},
		{
			given1:   GenericValue{"test", map[string]interface{}{"test": "test"}},
			given2:   GenericValue{"test", map[string]interface{}{"test": "test1"}},
			expected: false,
		},
		{
			given1:   GenericValue{"test", "1"},
			given2:   GenericValue{"test", 1},
			expected: true,
		},
	}

	for i, testCase := range testCases {
		result := testCase.given1.Equals(testCase.given2)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given1, "and", testCase.given2, "expected", testCase.expected, "result", result)
		}
	}
}

func TestVertexPropertyValueEquals(t *testing.T) {
	var testCases = []struct {
		given1   VertexPropertyValue
		given2   VertexPropertyValue
		expected bool
	}{
		{
			given1:   VertexPropertyValue{},
			given2:   VertexPropertyValue{},
			expected: true,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"},
			given2:   VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"},
			expected: true,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", 1}, "test", 1},
			given2:   VertexPropertyValue{GenericValue{"test", 1}, "test", 1},
			expected: true,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", 1}, "test", 1},
			given2:   VertexPropertyValue{GenericValue{"test", 2}, "test", 1},
			expected: false,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", 1}, "test", 1},
			given2:   VertexPropertyValue{GenericValue{"test", 1}, "test", 2},
			expected: false,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", false}, "test", false},
			given2:   VertexPropertyValue{GenericValue{"test", false}, "test", false},
			expected: true,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", false}, "test", false},
			given2:   VertexPropertyValue{GenericValue{"test", true}, "test", false},
			expected: false,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", false}, "test", false},
			given2:   VertexPropertyValue{GenericValue{"test", false}, "test", true},
			expected: false,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}},
			given2:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}},
			expected: true,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}},
			given2:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test1"}}, "test", map[string]interface{}{"test": "test"}},
			expected: false,
		},
		{
			given1:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}},
			given2:   VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test1"}},
			expected: false,
		},
	}

	for i, testCase := range testCases {
		result := testCase.given1.Equals(testCase.given2)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given1, "and", testCase.given2, "expected", testCase.expected, "result", result)
		}
	}
}

func TestVertexPropertyEquals(t *testing.T) {
	var testCases = []struct {
		given1   VertexProperty
		given2   VertexProperty
		expected bool
	}{
		{
			given1:   VertexProperty{},
			given2:   VertexProperty{},
			expected: true,
		},
		{
			given1:   VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
			given2:   VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
			expected: true,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			expected: true,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			given2:   VertexProperty{"test1", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 2}, "test", 1}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 2}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			expected: true,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			given2:   VertexProperty{"test1", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", false}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", false}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", false}, "test", true}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			expected: true,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			given2:   VertexProperty{"test1", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test1"}}, "test", map[string]interface{}{"test": "test"}}},
			expected: false,
		},
		{
			given1:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
			given2:   VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test1"}}},
			expected: false,
		},
	}

	for i, testCase := range testCases {
		result := testCase.given1.Equals(testCase.given2)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given1, "and", testCase.given2, "expected", testCase.expected, "result", result)
		}
	}
}

func TestVertexValueEquals(t *testing.T) {
	var testCases = []struct {
		given1   VertexValue
		given2   VertexValue
		expected bool
	}{
		{
			given1:   VertexValue{},
			given2:   VertexValue{},
			expected: true,
		},
		{
			given1:   VertexValue{ID: "test"},
			given2:   VertexValue{ID: "test"},
			expected: true,
		},
		{
			given1:   VertexValue{ID: "test"},
			given2:   VertexValue{ID: "test1"},
			expected: false,
		},
		{
			given1:   VertexValue{Label: "test"},
			given2:   VertexValue{Label: "test"},
			expected: true,
		},
		{
			given1:   VertexValue{Label: "test"},
			given2:   VertexValue{Label: "test1"},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: true,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test1": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test1"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 2}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", false}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test1"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test2": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test2": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: true,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test2": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test3": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
		{
			given1: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test2": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			given2: VertexValue{
				ID:    "test_id",
				Label: "test_label",
				Properties: map[string][]VertexProperty{
					"test": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
					},
					"test2": []VertexProperty{
						VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test5adfad", "test"}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
						VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
					},
				},
			},
			expected: false,
		},
	}

	for i, testCase := range testCases {
		result := testCase.given1.Equals(testCase.given2)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given1, "and", testCase.given2, "expected", testCase.expected, "result", result)
		}
	}
}

func TestVertexEquals(t *testing.T) {
	var testCases = []struct {
		given1   Vertex
		given2   Vertex
		expected bool
	}{
		{
			given1:   Vertex{},
			given2:   Vertex{},
			expected: true,
		},
		{
			given1:   Vertex{"type", VertexValue{}},
			given2:   Vertex{"type", VertexValue{}},
			expected: true,
		},
		{
			given1:   Vertex{"type", VertexValue{}},
			given2:   Vertex{"type2", VertexValue{}},
			expected: false,
		},
		{
			given1:   Vertex{"type", VertexValue{ID: "test"}},
			given2:   Vertex{"type", VertexValue{ID: "test"}},
			expected: true,
		},
		{
			given1:   Vertex{"type", VertexValue{ID: "test"}},
			given2:   Vertex{"type2", VertexValue{ID: "test"}},
			expected: false,
		},
		{
			given1:   Vertex{"type", VertexValue{ID: "test"}},
			given2:   Vertex{"type", VertexValue{ID: "test1"}},
			expected: false,
		},
		{
			given1:   Vertex{"type", VertexValue{Label: "test"}},
			given2:   Vertex{"type", VertexValue{Label: "test"}},
			expected: true,
		},
		{
			given1:   Vertex{"type", VertexValue{Label: "test"}},
			given2:   Vertex{"type", VertexValue{Label: "test1"}},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: true,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test1": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test1"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 2}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", false}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test1"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test2": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test2": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: true,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test2": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test3": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
		{
			given1: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test2": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test2", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			given2: Vertex{"type",
				VertexValue{
					ID:    "test_id",
					Label: "test_label",
					Properties: map[string][]VertexProperty{
						"test": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test", map[string]interface{}{"test": "test"}}},
						},
						"test2": []VertexProperty{
							VertexProperty{"", VertexPropertyValue{GenericValue{"test", "test"}, "test5adfad", "test"}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", 1}, "test2", 1}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", true}, "test2", true}},
							VertexProperty{"test", VertexPropertyValue{GenericValue{"test", map[string]interface{}{"test": "test"}}, "test2", map[string]interface{}{"test": "test"}}},
						},
					},
				},
			},
			expected: false,
		},
	}

	for i, testCase := range testCases {
		result := testCase.given1.Equals(testCase.given2)
		if result != testCase.expected {
			t.Error("test", i, "given", testCase.given1, "and", testCase.given2, "expected", testCase.expected, "result", result)
		}
	}
}
