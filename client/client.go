package client

import (
	"node_topology_discovery/model"
)

type Client interface {
	MakeRequest(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}
