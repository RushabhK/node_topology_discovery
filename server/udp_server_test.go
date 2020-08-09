package server

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"node_topology_discovery/client"
	"node_topology_discovery/mocks"
	"node_topology_discovery/model"
	"sync"
	"testing"
)

type UdpServerTestSuite struct {
	suite.Suite
	mockCtrl         *gomock.Controller
	discoveryService *mocks.MockNodesDiscoveryService
}

func TestUdpServerTestSuite(t *testing.T) {
	suite.Run(t, new(UdpServerTestSuite))
}

func (suite *UdpServerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.discoveryService = mocks.NewMockNodesDiscoveryService(suite.mockCtrl)
}

func (suite UdpServerTestSuite) TestShouldServeOneRequestAndStopTheServerWhenNodesCountIsTwo() {
	nodesCount := make(chan int, 1)
	nodesCount <- 2
	var wg sync.WaitGroup
	wg.Add(1)

	server := NewUdpServer(suite.discoveryService)
	port := "30001"
	go server.Serve(port, nodesCount, &wg)
	client := client.NewUdpClient(5, "localhost", port)
	request := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:30001"},
		DebugTrace:   "machine-1",
	}
	expectedResponse := model.NodesDiscoveryResponse{
		"localhost:30001": model.NodeInfo{
			Name:      "machine-1",
			IpAddress: "localhost",
			Port:      "30001",
			Neighbors: []model.NeighborInfo{
				{Port: "30002", IpAddress: "localhost"},
				{Port: "30003", IpAddress: "localhost"},
			}},
	}

	suite.discoveryService.EXPECT().Discover(request).Return(expectedResponse, nil)

	response, err := client.MakeRequest(request)

	wg.Wait()
	suite.Nil(err)
	suite.Equal(expectedResponse, response)
}

func (suite UdpServerTestSuite) TestShouldServeTwoRequestsAndStopTheServerWhenNodesCountIsThree() {
	nodesCount := make(chan int, 1)
	nodesCount <- 3
	var wg sync.WaitGroup
	wg.Add(1)

	server := NewUdpServer(suite.discoveryService)
	port := "30002"
	go server.Serve(port, nodesCount, &wg)
	client := client.NewUdpClient(500, "localhost", port)

	request1 := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:30001"},
		DebugTrace:   "machine-1",
	}
	expectedResponse1 := model.NodesDiscoveryResponse{
		"localhost:30001": model.NodeInfo{
			Name:      "machine-1",
			IpAddress: "localhost",
			Port:      "30001",
			Neighbors: []model.NeighborInfo{
				{Port: "30002", IpAddress: "localhost"},
				{Port: "30003", IpAddress: "localhost"},
			}},
	}

	suite.discoveryService.EXPECT().Discover(request1).Return(expectedResponse1, nil)

	response1, err1 := client.MakeRequest(request1)
	suite.Nil(err1)
	suite.Equal(expectedResponse1, response1)

	request2 := model.NodesDiscoveryRequest{
		VisitedNodes: []string{"localhost:30001", "localhost:30002"},
		DebugTrace:   "machine-1 -> machine-2",
	}
	expectedResponse2 := model.NodesDiscoveryResponse{
		"localhost:30002": model.NodeInfo{
			Name:      "machine-2",
			IpAddress: "localhost",
			Port:      "30002",
			Neighbors: []model.NeighborInfo{
				{Port: "30004", IpAddress: "localhost"},
				{Port: "30003", IpAddress: "localhost"},
				{Port: "30002", IpAddress: "localhost"},
			}},
	}

	suite.discoveryService.EXPECT().Discover(request2).Return(expectedResponse2, nil)

	response2, err2 := client.MakeRequest(request2)

	suite.Nil(err2)
	suite.Equal(expectedResponse2, response2)
	wg.Wait()
}
