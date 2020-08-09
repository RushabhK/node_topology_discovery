package udp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/model"
	"time"
)

type Client interface {
	MakeRequest(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}

type udpClient struct {
	timeoutInSeconds time.Duration
	ipAddress        string
	port             string
}

func NewUdpClient(timeoutInSeconds time.Duration, ipAddress, port string) Client {
	return udpClient{
		timeoutInSeconds: timeoutInSeconds,
		port:             port,
		ipAddress:        ipAddress,
	}
}

func (udpClient udpClient) MakeRequest(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	addressString := udpClient.ipAddress + ":" + udpClient.port
	fmt.Println("Dialing udp..")
	udpConnection, err := net.Dial("udp", addressString)

	if err != nil {
		fmt.Println(err)
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

	responseBytes := make([]byte, 2048)
	
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
