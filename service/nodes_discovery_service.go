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
	client     client.Client
}

func NewNodesDiscoveryService(configData config_loader.ConfigData, client client.Client) NodesDiscoveryService {
	return nodesDiscoveryService{
		configData: configData,
		client:     client,
	}
}

func (service nodesDiscoveryService) Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	aggregatedResponse := model.NodesDiscoveryResponse{}
	aggregatedResponse[service.configData.GetIdentifier()] = service.configData.ToNodeInfo()
	request.UpdateVisitedNodes(aggregatedResponse)
	for _, neighbor := range service.configData.Neighbors {
		if !visited(neighbor, request.VisitedNodes) {
			discoveryResponse, err := service.client.MakeRequest(request)
			if err != nil {
				return model.NodesDiscoveryResponse{}, err
			}
			request.UpdateVisitedNodes(discoveryResponse)
			aggregatedResponse = aggregatedResponse.Union(discoveryResponse)
		}
	}
	return aggregatedResponse, nil
}

func visited(neighbor config_loader.Neighbor, visitedNodes []string) bool {
	for _, visitedNode := range visitedNodes {
		if neighbor.GetIdentifier() == visitedNode {
			return true
		}
	}
	return false
}
