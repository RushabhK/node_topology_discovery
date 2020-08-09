package main

import (
	"fmt"
	"node_topology_discovery/config_loader"
	"node_topology_discovery/constants"
	"node_topology_discovery/model"
	"node_topology_discovery/service"
	"node_topology_discovery/tcp"
	"sync"
)

func main() {
	configLoader := config_loader.NewConfigLoader()
	configData, loadError := configLoader.Load(constants.CONFIG_FILE_PATH)

	if loadError != nil {
		panic("Cannot load config: " + loadError.Error())
	}

	discoveryService := service.NewNodesDiscoveryService(configData)
	tcpServer := tcp.NewTcpServer(discoveryService)

	var wg sync.WaitGroup
	wg.Add(1)
	nodesCount := make(chan int, 1)

	go tcpServer.Serve(configData.Port, nodesCount, &wg)

	discoveryRequest := model.NodesDiscoveryRequest{
		DebugTrace:   configData.Name,
		VisitedNodes: []string{configData.GetIdentifier()},
	}
	discoveryResponse, discoveryError := discoveryService.Discover(discoveryRequest)
	if discoveryError != nil {
		fmt.Println("Error while discovering: ", discoveryError.Error())
		nodesCount <- 1
	}
	nodesCount <- len(discoveryResponse)
	wg.Wait()
}

/*
E - A - B
	|   |
F - C - D

1. Tcp Server running on A
	- ReqCount = 0

2. Discovery service -> Discover all nodes in topology
	- 2nd goroutine
	- { VisitedNodes -> [A] }

Req to E is { VisitedNodes -> [A] }: Response from E is : {E: [A]}
Req to B is { VisitedNodes -> [A,E] }: Response from E is : {E: [A]}

	- Response:
{
	B : [A, D],
	C : [A, D],
	D : [C, B],
	E : [A],
	A : [E, C, B]
}
 */