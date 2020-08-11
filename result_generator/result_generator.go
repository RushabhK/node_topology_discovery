//TODO: Add unit tests and refactor result generator
package result_generator

import (
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
	Generate(nodesCount chan int, wg *sync.WaitGroup, machineName string) error
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

func (resultGenerator resultGenerator) Generate(nodesCount chan int, wg *sync.WaitGroup, machineName string) error {
	defer func() {
		wg.Done()
	}()

	var finalResult string
	var executionTime float64

	defer func() {
		fmt.Println("FINAL RESULT: ", finalResult)
		fmt.Println("EXECUTION TIME: ", executionTime)
		executionTimeFormat := fmt.Sprintf("Total time taken for discovery for machine %s: %f seconds\n", machineName, executionTime)
		fileContent := finalResult + "\n" + executionTimeFormat
		resultGenerator.fileUtils.Write(constants.RESULT_FILE_PATH, fileContent)
	}()

	fmt.Println("Discovering Starts.....")
	start := time.Now()
	discoveryResponse, err := resultGenerator.service.Discover(model.NodesDiscoveryRequest{})
	elapsed := time.Since(start)
	executionTime = elapsed.Seconds()

	if err != nil {
		fmt.Println("Error encountered while discovering:", err.Error())
		fmt.Println("Writing to NODES_COUNT CHANNEL value : 1")
		nodesCount <- 1
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

	finalResult = formatContent(resultMap, machineName)
	fmt.Println("Writing to NODES_COUNT CHANNEL value : ", len(resultMap))
	nodesCount <- len(resultMap)
	return nil
}

func formatContent(resultMap map[string][]string, machineName string) string {
	result := fmt.Sprintf("There are %v machines in this topology. The following are the machines and their neighbors discovered from machine %s:\n", len(resultMap), machineName)
	keys := make([]string, 0, len(resultMap))
	for k := range resultMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		result += fmt.Sprintf("%s: [", k)
		v := resultMap[k]
		for _, n := range v {
			result += n + ", "
		}
		result = result[:len(result)-2] + "]\n"
	}
	return result
}
