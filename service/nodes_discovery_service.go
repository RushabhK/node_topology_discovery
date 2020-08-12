package service

import (
	"fmt"
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

func NewNodesDiscoveryService(configData config_loader.ConfigData, clientFactory client.ClientFactory) NodesDiscoveryService {
	return nodesDiscoveryService{
		configData: configData,
		client:     clientFactory.GetClient(),
	}
}

func (service nodesDiscoveryService) Discover(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	//request.DebugTrace += service.configData.Name + " -> "
	fmt.Println("Discovering the machines in path: ", request.DebugTrace)
	aggregatedResponse := model.NodesDiscoveryResponse{}
	aggregatedResponse[service.configData.GetIdentifier()] = service.configData.ToNodeInfo()
	request.UpdateVisitedNodes(aggregatedResponse)
	for _, neighbor := range service.configData.Neighbors {
		if !visited(neighbor, request.VisitedNodes) {
			discoveryResponse, err := service.client.MakeRequest(neighbor.IpAddress, neighbor.Port, request)
			if err != nil {
				fmt.Println("Error while discovering: ", err.Error())
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
