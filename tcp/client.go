package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/model"
	"time"
)

type Client interface {
	MakeRequest(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error)
}

type tcpClient struct {
	timeoutInSeconds time.Duration
	ipAddress        string
	port             string
}

func NewTcpClient(timeoutInSeconds time.Duration, ipAddress, port string) Client {
	return tcpClient{
		timeoutInSeconds: timeoutInSeconds,
		port:             port,
		ipAddress:        ipAddress,
	}
}

func (tcpClient tcpClient) MakeRequest(request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	addressString := tcpClient.ipAddress + ":" + tcpClient.port
	dialer := net.Dialer{Timeout: tcpClient.timeoutInSeconds * time.Second}
	tcpConnection, err := dialer.Dial("tcp", addressString)
	//defer bufio.NewWriter(tcpConnection).Flush()

	if err != nil {
		fmt.Println(err)
		return model.NodesDiscoveryResponse{}, err
	}

	requestBytes, _ := json.Marshal(request)
	requestString := string(requestBytes)
	fmt.Fprintf(tcpConnection, requestString + "\n")

	decoder := json.NewDecoder(tcpConnection)
	var response model.NodesDiscoveryResponse
	decodeErr := decoder.Decode(&response)

	/*responseString, readErr := bufio.NewReader(tcpConnection).ReadString('\n')
	if readErr != nil {
		return model.NodesDiscoveryResponse{}, readErr
	}
	var response model.NodesDiscoveryResponse
	unmarshalErr := json.Unmarshal([]byte(responseString), &request)*/

	return response, decodeErr
}
