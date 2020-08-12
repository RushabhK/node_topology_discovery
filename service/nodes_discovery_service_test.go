package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"node_topology_discovery/config_loader"
	"node_topology_discovery/mocks"
	"node_topology_discovery/model"
	"testing"
)

type NodesDiscoveryServiceTestSuite struct {
	suite.Suite
	mockCtrl      *gomock.Controller
	clientFactory *mocks.MockClientFactory
	client        *mocks.MockClient
}

func TestNodesDiscoveryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NodesDiscoveryServiceTestSuite))
}

func (suite *NodesDiscoveryServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.clientFactory = mocks.NewMockClientFactory(suite.mockCtrl)
	suite.client = mocks.NewMockClient(suite.mockCtrl)
	suite.clientFactory.EXPECT().GetClient().Return(suite.client)
}

func (suite NodesDiscoveryServiceTestSuite) TestShouldDiscoverAllNodesAndGiveAggregatedResponse() {
	configData := config_loader.ConfigData{
		IpAddress: "localhost",
		Port:      "3202",
		Name:      "machine-2",
		Neighbors: []config_loader.Neighbor{
			{IpAddress: "localhost", Port: "3203"},
			{IpAddress: "localhost", Port: "3204"},
		},
	}

	discoveryService := NewNodesDiscoveryService(configData, suite.clientFactory)
	request := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201"},
	}
	request3 := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201", "localhost:3202"},
	}
	response3 := model.NodesDiscoveryResponse{
		"localhost:3203": model.NodeInfo{
			Name:      "machine-3",
			IpAddress: "localhost",
			Port:      "3203",
			Neighbors: []model.NeighborInfo{{IpAddress: "localhost", Port: "3202"}},
		},
	}
	suite.client.EXPECT().MakeRequest("localhost", "3203", request3).Return(response3, nil)

	request4 := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201", "localhost:3202", "localhost:3203"},
	}
	response4 := model.NodesDiscoveryResponse{
		"localhost:3204": model.NodeInfo{
			Name:      "machine-4",
			IpAddress: "localhost",
			Port:      "3204",
			Neighbors: []model.NeighborInfo{{IpAddress: "localhost", Port: "3202"}},
		},
	}
	suite.client.EXPECT().MakeRequest("localhost", "3204", request4).Return(response4, nil)
	response, discoverErr := discoveryService.Discover(request)

	suite.Nil(discoverErr)

	expectedResponse := model.NodesDiscoveryResponse{
		"localhost:3202": model.NodeInfo{
			Name:      "machine-2",
			IpAddress: "localhost",
			Port:      "3202",
			Neighbors: []model.NeighborInfo{
				{IpAddress: "localhost", Port: "3203"},
				{IpAddress: "localhost", Port: "3204"},
			},
		},
		"localhost:3203": model.NodeInfo{
			Name:      "machine-3",
			IpAddress: "localhost",
			Port:      "3203",
			Neighbors: []model.NeighborInfo{{IpAddress: "localhost", Port: "3202"}},
		},
		"localhost:3204": model.NodeInfo{
			Name:      "machine-4",
			IpAddress: "localhost",
			Port:      "3204",
			Neighbors: []model.NeighborInfo{{IpAddress: "localhost", Port: "3202"}},
		},
	}

	suite.Equal(expectedResponse, response)
}

func (suite NodesDiscoveryServiceTestSuite) TestShouldNotDiscoverNodesIfAlreadyVisited() {
	configData := config_loader.ConfigData{
		IpAddress: "localhost",
		Port:      "3202",
		Name:      "machine-2",
		Neighbors: []config_loader.Neighbor{
			{IpAddress: "localhost", Port: "3203"},
			{IpAddress: "localhost", Port: "3204"},
		},
	}

	discoveryService := NewNodesDiscoveryService(configData, suite.clientFactory)
	request := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201", "localhost:3203", "localhost:3204"},
	}

	response, discoverErr := discoveryService.Discover(request)
	suite.Nil(discoverErr)

	expectedResponse := model.NodesDiscoveryResponse{
		"localhost:3202": model.NodeInfo{
			Name:      "machine-2",
			IpAddress: "localhost",
			Port:      "3202",
			Neighbors: []model.NeighborInfo{
				{IpAddress: "localhost", Port: "3203"},
				{IpAddress: "localhost", Port: "3204"},
			},
		},
	}

	suite.Equal(expectedResponse, response)
}

func (suite NodesDiscoveryServiceTestSuite) TestShouldReturnErrorIfDiscoveryOnAnyNeighborFails() {
	configData := config_loader.ConfigData{
		IpAddress: "localhost",
		Port:      "3202",
		Name:      "machine-2",
		Neighbors: []config_loader.Neighbor{
			{IpAddress: "localhost", Port: "3203"},
			{IpAddress: "localhost", Port: "3204"},
		},
	}

	discoveryService := NewNodesDiscoveryService(configData, suite.clientFactory)
	request := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201"},
	}

	request3 := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:3201", "localhost:3202"},
	}
	err := errors.New("Discovery failed on localhost:3203")
	suite.client.EXPECT().MakeRequest("localhost", "3203", request3).Return(model.NodesDiscoveryResponse{}, err)
	response, discoverErr := discoveryService.Discover(request)

	suite.Equal(err, discoverErr)
	suite.Empty(response)
}
