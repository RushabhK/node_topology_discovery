package tcp

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"node_topology_discovery/mocks"
	"node_topology_discovery/model"
	"sync"
	"testing"
)

type TcpClientServerTestSuite struct {
	suite.Suite
	mockCtrl         *gomock.Controller
	discoveryService *mocks.MockNodesDiscoveryService
}

func TestTcpServerTestSuite(t *testing.T) {
	suite.Run(t, new(TcpClientServerTestSuite))
}

func (suite *TcpClientServerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.discoveryService = mocks.NewMockNodesDiscoveryService(suite.mockCtrl)
}

func (suite TcpClientServerTestSuite) TestShouldServeOneRequestAndStopTheServerWhenNodesCountIsTwo() {
	nodesCount := make(chan int, 1)
	nodesCount <- 2
	var wg sync.WaitGroup
	wg.Add(1)

	server := NewTcpServer(suite.discoveryService)
	port := "30001"
	go server.Serve(port, nodesCount, &wg)
	client := NewTcpClient(5, "localhost", port)
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
