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
	println("Dialing udp")
	udpConnection, err := net.Dial("udp", addressString)

	if err != nil {
		fmt.Println(err)
		return model.NodesDiscoveryResponse{}, err
	}
	defer udpConnection.Close()

	requestBytes, _ := json.Marshal(request)
	//requestString := string(requestBytes)
	udpConnection.Write(requestBytes)
	//fmt.Fprintf(udpConnection, requestString + "\n")

	p := make([]byte, 2048)
	readLen, err := bufio.NewReader(udpConnection).Read(p)
	if err != nil {
		return model.NodesDiscoveryResponse{}, err
	}
	var response model.NodesDiscoveryResponse
	unmarshalErr := json.Unmarshal(p[:readLen], &response)

	return response, unmarshalErr
}
