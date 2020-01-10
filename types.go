package gremlin

type GremlinQuery struct {
	Query   string
	Args    []interface{}
	LockKey string
}

// cbi made up, not a real graphson or gremlin thing
type GremlinResponse struct {
	V VertexesV2
	E Edges
}

type VertexesV2 []VertexV2

type Vertexes []Vertex

type VertexV2 struct {
	Type  string        `json:"@type"`
	Value VertexValueV2 `json:"@value"`
}

type Vertex struct {
	Type  string      `json:"@type"`
	Value VertexValue `json:"@value"`
}

func (v1 Vertex) Equals(v2 Vertex) bool {
	return v1.Type == v2.Type && v1.Value.Equals(v2.Value)
}

type IdTypeStruct struct {
	Type  string      `json:"@type"`
	Value int `json:"@value"`
}

type IdType IdTypeStruct

type VertexValue struct {
	ID         string                      `json:"id"`
	Label      string                      `json:"label"`
	Properties map[string][]VertexProperty `json:"properties"`
}

type VertexValueV2 struct {
	ID         IdType                      `json:"id"`
	Label      string                      `json:"label"`
	Properties map[string][]VertexProperty `json:"properties"`
}

func (v1 VertexValue) Equals(v2 VertexValue) bool {
	if v1.ID != v2.ID || v1.Label != v2.Label || len(v1.Properties) != len(v2.Properties) {
		return false
	}
	for k, val1Slice := range v1.Properties {
		val2Slice, ok := v2.Properties[k]
		if !ok || len(val1Slice) != len(val2Slice) {
			return false
		}
		for i, val1 := range val1Slice {
			val2 := val2Slice[i]
			if !val1.Equals(val2) {
				return false
			}
		}
	}
	return true
}

type VertexProperty struct {
	Type  string              `json:"@type"`
	Value VertexPropertyValue `json:"@value"`
}

func (v1 VertexProperty) Equals(v2 VertexProperty) bool {
	return v1.Type == v2.Type && v1.Value.Equals(v2.Value)
}

type EdgeProperty struct {
	Type  string            `json:"@type"`
	Value EdgePropertyValue `json:"@value"`
}

type VertexPropertyValue struct {
	ID    GenericValue `json:"id"`
	Label string       `json:"label"`
	Value interface{}  `json:"value"`
}

func (v1 VertexPropertyValue) Equals(v2 VertexPropertyValue) bool {
	return v1.ID.Equals(v2.ID) && v1.Label == v2.Label && InterfacesMatch(v1.Value, v2.Value)
}

type EdgePropertyValue struct {
	Label string      `json:"key"`
	Value interface{} `json:"value"`
}

type GenericValues []GenericValue

type GenericValue struct {
	Type  string      `json:"@type"`
	Value interface{} `json:"@value"`
}

func (g1 GenericValue) Equals(g2 GenericValue) bool {
	return g1.Type == g2.Type && InterfacesMatch(g1.Value, g2.Value)
}

type Edges []Edge
type EdgesV2 []EdgeV2

type Edge struct {
	Type  string    `json:"@type"`
	Value EdgeValue `json:"@value"`
}

type EdgeV2 struct {
	Type  string    `json:"@type"`
	Value EdgeValueV2 `json:"@value"`
}

type EdgeIdValueType struct {
	RelationId  string
}

type EdIdType struct {
	Type  string    `json:"@type"`
	Value EdgeIdValueType `json:"@value"`
}

type EdgeValue struct {
	ID         string // TODO: does this need to be a GenericValue? interface{}?
	Label      string
	InVLabel   string
	OutVLabel  string
	InV        string
	OutV       string
	Properties map[string]EdgeProperty
}

type EdgeValueV2 struct {
	ID		   EdIdType
	Label      string
	InVLabel   string
	OutVLabel  string
	InV        string
	OutV       string
	Properties map[string]EdgeProperty
}

type CleanResponse struct {
	V []CleanVertex
	E []CleanEdge
}

type CleanEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type CleanVertex struct {
	Id    int `json:"id"`
	Label string `json:"label"`
}

// type TypeID int

// const (
// 	TypeString TypeID = iota
// 	TypeBoolean
// 	TypeMap
// 	TypeCollection
// 	TypeClass
// 	TypeDate
// 	TypeDouble
// 	TypeFloat
// 	TypeInteger
// 	TypeLong
// 	TypeTimestamp
// 	TypeUUID
// 	TypeVertex
// 	TypeVertexProperty
// )

// const (
// 	TypeStrDate           = "g:Date"
// 	TypeStrDouble         = "g:Double"
// 	TypeStrFloat          = "g:Float"
// 	TypeStrInteger        = "g:Int32"
// 	TypeStrLong           = "g:Int64"
// 	TypeStrTimestamp      = "g:Timestamp"
// 	TypeStrUUID           = "g:UUID"
// 	TypeStrVertex         = "g:Vertex"
// 	TypeStrVertexProperty = "g:VertexProperty"
// 	TypeStrProperty       = "g:Property"
// 	TypeStrEdge           = "g:Edge"
// )
