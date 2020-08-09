package service

import (
	"node_topology_discovery/config_loader"
	"node_topology_discovery/model"
)

type NodesDiscoveryService interface {
	Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}

type nodesDiscoveryService struct {
	configData config_loader.ConfigData
}

func NewNodesDiscoveryService(configData config_loader.ConfigData) NodesDiscoveryService {
	return nodesDiscoveryService{
		configData: configData,
	}
}

func (nodesDiscoveryService) Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	panic("implement me")
}
