package config_loader

import "node_topology_discovery/model"

type ConfigData struct {
	Name      string     `json:"name"`
	IpAddress string     `json:"ip_address"`
	Port      string     `json:"port"`
	Neighbors []Neighbor `json:"neighbors"`
}

type Neighbor struct {
	IpAddress string `json:"ip_address"`
	Port      string `json:"port"`
}

func (neighbor Neighbor) ToNeighborInfo() model.NeighborInfo {
	return model.NeighborInfo{
		IpAddress: neighbor.IpAddress,
		Port:      neighbor.Port,
	}
}

func (configData ConfigData) GetIdentifier() string {
	return configData.IpAddress + ":" + configData.Port
}

func (neighbor Neighbor) GetIdentifier() string {
	return neighbor.IpAddress + ":" + neighbor.Port
}

func (configData ConfigData) ToNodeInfo() model.NodeInfo {
	nodeInfo := model.NodeInfo{
		Name:      configData.Name,
		Port:      configData.Port,
		IpAddress: configData.IpAddress,
	}

	for _, neighbor := range configData.Neighbors {
		nodeInfo.Neighbors = append(nodeInfo.Neighbors, neighbor.ToNeighborInfo())
	}

	return nodeInfo
}
