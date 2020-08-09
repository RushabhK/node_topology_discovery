package model

import "encoding/json"

type NodesDiscoveryRequest struct {
	VisitedNodes []string `json:"visited_nodes"`
	DebugTrace   string   `json:"debug_trace"`
}

func (request NodesDiscoveryRequest) ToString() string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}
