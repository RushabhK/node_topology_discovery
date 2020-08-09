package model

import "encoding/json"

type NodeInfo struct {
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Port      string `json:"port"`
	Neighbors []NeighborInfo
}

type NeighborInfo struct {
	IpAddress string `json:"ip_address"`
	Port      string `json:"port"`
}

func (nodeInfo NodeInfo) GetIdentifier() string {
	return nodeInfo.IpAddress + ":" + nodeInfo.Port
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
