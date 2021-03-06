package main

import (
	"node_topology_discovery/client"
	"node_topology_discovery/config_loader"
	"node_topology_discovery/constants"
	"node_topology_discovery/result_generator"
	"node_topology_discovery/server"
	"node_topology_discovery/service"
	"node_topology_discovery/utils"
	"sync"
	"time"
)

func main() {
	configLoader := config_loader.NewConfigLoader()
	configData, loadError := configLoader.Load(constants.CONFIG_FILE_PATH)

	if loadError != nil {
		panic("Cannot load config: " + loadError.Error())
	}

	udpClient := client.NewUdpClientFactory(constants.CLIENT_DIAL_TIMEOUT)
	discoveryService := service.NewNodesDiscoveryService(configData, udpClient)
	server := server.NewUdpServer(discoveryService)
	fileUtils := utils.NewFileUtils()
	resultGenerator := result_generator.NewResultGenerator(discoveryService, fileUtils)

	var wg sync.WaitGroup
	wg.Add(2)
	nodesCount := make(chan int, 1)

	go server.Serve(configData.Port, nodesCount, &wg)
	time.Sleep(constants.INITIAL_WAIT_IN_SEC * time.Second)
	go resultGenerator.Generate(nodesCount, &wg, configData.Name)

	wg.Wait()
}
