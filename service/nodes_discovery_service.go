package service

import (
	"node_topology_discovery/client"
	"node_topology_discovery/config_loader"
	"node_topology_discovery/model"
)

type NodesDiscoveryService interface {
	Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}

type nodesDiscoveryService struct {
	configData config_loader.ConfigData
	udpClient  client.Client
}

func NewNodesDiscoveryService(configData config_loader.ConfigData) NodesDiscoveryService {
	return nodesDiscoveryService{
		configData: configData,
	}
}

func (service nodesDiscoveryService) Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	response := model.NodesDiscoveryResponse{}
	response[service.configData.GetIdentifier()] = service.configData.ToNodeInfo()
	panic("")
}
