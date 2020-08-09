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

func (request *NodesDiscoveryRequest) UpdateVisitedNodes(response NodesDiscoveryResponse) {
	for newVisitedNode := range response {
		if !visited(request.VisitedNodes, newVisitedNode) {
			request.VisitedNodes = append(request.VisitedNodes, newVisitedNode)
		}
	}
}

func visited(visitedNodes []string, newVisitedNode string) bool {
	for _, visitedNode := range visitedNodes {
		if newVisitedNode == visitedNode {
			return true
		}
	}
	return false
}
