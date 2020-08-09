package model

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type NodesDiscoveryResponseTestSuite struct {
	suite.Suite
}

func TestNodesDiscoveryResponseTestSuite(t *testing.T) {
	suite.Run(t, new(NodesDiscoveryResponseTestSuite))
}

func (suite NodesDiscoveryResponseTestSuite) TestShouldMakeUnionOfTwoResponses() {
	response1 := NodesDiscoveryResponse{
		"localhost:3200": NodeInfo{
			IpAddress: "localhost",
			Port:      "3200",
			Name:      "machine-0",
			Neighbors: []NeighborInfo{{Port: "3201", IpAddress: "localhost"}},
		},
	}
	response2 := NodesDiscoveryResponse{
		"localhost:3202": NodeInfo{
			IpAddress: "localhost",
			Port:      "3202",
			Name:      "machine-2",
			Neighbors: []NeighborInfo{{Port: "3203", IpAddress: "localhost"}},
		},
	}

	expectedUnion := NodesDiscoveryResponse{
		"localhost:3200": NodeInfo{
			IpAddress: "localhost",
			Port:      "3200",
			Name:      "machine-0",
			Neighbors: []NeighborInfo{{Port: "3201", IpAddress: "localhost"}},
		},
		"localhost:3202": NodeInfo{
			IpAddress: "localhost",
			Port:      "3202",
			Name:      "machine-2",
			Neighbors: []NeighborInfo{{Port: "3203", IpAddress: "localhost"}},
		},
	}

	suite.Equal(expectedUnion, response1.Union(response2))

}
