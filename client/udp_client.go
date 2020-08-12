package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/constants"
	"node_topology_discovery/model"
	"time"
)

type udpClient struct {
	timeoutInSeconds time.Duration
}

func (udpClient udpClient) MakeRequest(ipAddress string, port string, request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	addressString := ipAddress + ":" + port
	fmt.Println("Dialing udp..")
	dialer := net.Dialer{Timeout: udpClient.timeoutInSeconds * time.Second}
	udpConnection, err := dialer.Dial("udp", addressString)

	if err != nil {
		fmt.Println("Error while dialing: ", err.Error())
		return model.NodesDiscoveryResponse{}, err
	}
	defer udpConnection.Close()

	requestBytes, _ := json.Marshal(request)
	fmt.Println("Writing to UDP connection : ", string(requestBytes))
	_, writeErr := udpConnection.Write(requestBytes)
	if writeErr != nil {
		fmt.Println("Error while writing : ", writeErr.Error())
		return model.NodesDiscoveryResponse{}, writeErr
	}
	fmt.Println("Write successful!")

	responseBytes := make([]byte, constants.BUFFER_SIZE)

	fmt.Println("Reading response from udp server... ")
	readLen, err := bufio.NewReader(udpConnection).Read(responseBytes)
	if err != nil {
		fmt.Println("Error while reading response from udp server : ", err.Error())
		return model.NodesDiscoveryResponse{}, err
	}
	fmt.Println("Response from udp server : ", string(responseBytes))
	var response model.NodesDiscoveryResponse
	fmt.Println("Deserializing the response..")
	unmarshalErr := json.Unmarshal(responseBytes[:readLen], &response)
	if unmarshalErr != nil {
		fmt.Println("Error while deserilizing the response : ", unmarshalErr.Error())
		return response, unmarshalErr
	}
	fmt.Println("Deserialization of response done successfully")
	return response, nil
}
