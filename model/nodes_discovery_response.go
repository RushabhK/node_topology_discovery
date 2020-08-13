package model

import "encoding/json"

type NodeInfo struct {
	Name      string         `json:"n"`
	IpAddress string         `json:"i"`
	Port      string         `json:"p"`
	Neighbors []NeighborInfo `json:"N"`
}

type NeighborInfo struct {
	IpAddress string `json:"i"`
	Port      string `json:"p"`
}

func (nodeInfo NodeInfo) GetIdentifier() string {
	return nodeInfo.IpAddress + ":" + nodeInfo.Port
}

func (neighborInfo NeighborInfo) GetIdentifier() string {
	return neighborInfo.IpAddress + ":" + neighborInfo.Port
}

type NodesDiscoveryResponse map[string]NodeInfo

func (response NodesDiscoveryResponse) ToString() string {
	bytes, _ := json.Marshal(response)
	return string(bytes)
}

func (response1 NodesDiscoveryResponse) Union(response2 NodesDiscoveryResponse) NodesDiscoveryResponse {
	for k, v := range response2 {
		response1[k] = v
	}
	return response1
}
