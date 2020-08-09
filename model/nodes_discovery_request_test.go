package model

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type NodesDiscoveryRequestTestSuite struct {
	suite.Suite
}

func TestNodesDiscoveryRequestTestSuite(t *testing.T) {
	suite.Run(t, new(NodesDiscoveryRequestTestSuite))
}

func (suite NodesDiscoveryRequestTestSuite) TestShouldMakeUnionOfTwoResponses() {
	request := NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3200", "localhost:3201"},
	}
	response := NodesDiscoveryResponse{
		"localhost:3200": NodeInfo{},
		"localhost:3201": NodeInfo{},
		"localhost:3202": NodeInfo{},
		"localhost:3203": NodeInfo{},
	}

	request.UpdateVisitedNodes(response)

	expectedRequest := NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3200", "localhost:3201", "localhost:3202", "localhost:3203"},
	}
	suite.Equal(expectedRequest, request)

}
