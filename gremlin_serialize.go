package gremlin

import (
	"encoding/json"
	"strconv"
)

func DeserializeVertices(rawResponse string) (Vertexes, error) {
	var response Vertexes
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DeserializeVerticesV2(rawResponse string) (VertexesV2, error) {
	// TODO: empty strings for property values will cause invalid json
	// make so it can handle that case
	var response VertexesV2
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DeserializeEdges(rawResponse string) (Edges, error) {
	var response Edges
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DeserializeEdgesV2(rawResponse string) (EdgesV2, error) {
	var response EdgesV2
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func SerializeGenericValues(rawResponse string) (GenericValues, error) {
	var response GenericValues
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func SerializeListInterface(rawResponse string) ([]interface{}, error) {
	var response []interface{}
	if rawResponse == "" {
		return response, nil
	}
	err := json.Unmarshal([]byte(rawResponse), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ConvertToCleanVertexes(vertexes Vertexes) []CleanVertex {
	var responseVertexes []CleanVertex
	for _, vertex := range vertexes {
		idAsInt, err := strconv.Atoi(vertex.Value.ID)
		if err != nil {
			panic(err)
		}
		responseVertexes = append(responseVertexes, CleanVertex{
			Id:    idAsInt,
			Label: vertex.Value.Label,
		})
	}
	return responseVertexes
}

func ConvertToCleanEdges(edges Edges) []CleanEdge {
	var responseEdges []CleanEdge
	for _, edge := range edges {
		responseEdges = append(responseEdges, CleanEdge{
			Source: edge.Value.InV,
			Target: edge.Value.OutV,
		})
	}
	return responseEdges
}
