//TODO: Add unit tests and refactor result generator
package result_generator

import (
	"encoding/json"
	"fmt"
	"node_topology_discovery/constants"
	"node_topology_discovery/model"
	"node_topology_discovery/service"
	"node_topology_discovery/utils"
	"sort"
	"sync"
	"time"
)

type ResultGenerator interface {
	Generate(nodesCount chan int, wg *sync.WaitGroup) error
}

type resultGenerator struct {
	service   service.NodesDiscoveryService
	fileUtils utils.FileUtils
}

func NewResultGenerator(service service.NodesDiscoveryService, fileUtils utils.FileUtils) ResultGenerator {
	return resultGenerator{
		service:   service,
		fileUtils: fileUtils,
	}
}

func (resultGenerator resultGenerator) Generate(nodesCount chan int, wg *sync.WaitGroup) error {
	defer func() {
		wg.Done()
	}()

	var finalResult string
	var executionTime float64

	defer func() {
		fmt.Println("FINAL RESULT: ", finalResult)
		fmt.Println("EXECUTION TIME: ", executionTime)
		fileContent := "FINAL RESULT: \n" + finalResult + "\n\nEXECUTION TIME: " + fmt.Sprintf("%f seconds\n", executionTime)
		resultGenerator.fileUtils.Write(constants.RESULT_FILE_PATH, fileContent)
	}()

	fmt.Println("Discovering Starts.....")
	start := time.Now()
	discoveryResponse, err := resultGenerator.service.Discover(model.NodesDiscoveryRequest{})
	elapsed := time.Since(start)
	executionTime = elapsed.Seconds()

	if err != nil {
		fmt.Println("Error encountered while discovering")
		fmt.Println("Writing to NODES_COUNT CHANNEL value : 1")
		nodesCount <- 1
		finalResult = err.Error()
		return err
	}

	machineMap := make(map[string]string)
	for k, v := range discoveryResponse {
		machineMap[k] = v.Name
	}

	resultMap := make(map[string][]string)
	for k, v := range discoveryResponse {
		machine := machineMap[k]
		for _, neighbor := range v.Neighbors {
			neighborName := machineMap[neighbor.GetIdentifier()]
			resultMap[machine] = append(resultMap[machine], neighborName)
		}
		sort.Strings(resultMap[machine])
	}

	bytes, _ := json.Marshal(resultMap)
	finalResult = string(bytes)
	fmt.Println("Writing to NODES_COUNT CHANNEL value : ", len(resultMap))
	nodesCount <- len(resultMap)
	return nil
}
