package client

import (
	"node_topology_discovery/model"
)

type Client interface {
	MakeRequest(ipAddress string, port string, request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}
