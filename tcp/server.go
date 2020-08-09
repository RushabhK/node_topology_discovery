package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/model"
	"node_topology_discovery/service"
	"sync"
)

type Server interface {
	Serve(port string, nodesCount chan int, wg *sync.WaitGroup) error
}

type tcpServer struct {
	requestsServed  int
	totalNodesCount int
	service         service.NodesDiscoveryService
}

func NewTcpServer(discoveryService service.NodesDiscoveryService) Server {
	return tcpServer{
		service:         discoveryService,
		requestsServed:  0,
		totalNodesCount: 0,
	}
}

func (server tcpServer) Serve(port string, nodesCount chan int, wg *sync.WaitGroup) (serveError error) {
	defer wg.Done()

	portAddress := ":" + port
	listener, listenError := net.Listen("tcp", portAddress)
	if listenError != nil {
		fmt.Println("Error while listening : ", listenError.Error())
		serveError = listenError
	}

	defer listener.Close()

	tcpConnection, acceptError := listener.Accept()
	if acceptError != nil {
		fmt.Println("Error while accepting tcp connections : ", acceptError.Error())
		return acceptError
	}

	for {
		if len(nodesCount) > 0 {
			server.totalNodesCount = <- nodesCount
		}
		if server.totalNodesCount != 0 && server.requestsServed == server.totalNodesCount - 1 {
			return nil
		}

		decoder := json.NewDecoder(tcpConnection)
		var request model.NodesDiscoveryRequest
		decodeErr := decoder.Decode(&request)
		if decodeErr != nil {
			serveError = decodeErr
		}

		/*requestString, readErr := bufio.NewReader(tcpConnection).ReadString('\n')
		if readErr != nil {
			serveError = readErr
		}
		var request model.NodesDiscoveryRequest
		unmarshalErr := json.Unmarshal([]byte(requestString), &request)
		if unmarshalErr != nil {
			serveError = unmarshalErr
		}*/

		fmt.Println("Request received : ", request.ToString())

		if len(request.VisitedNodes) > 0 {
			discoveryResponse, responseErr := server.service.Discover(request)
			if responseErr != nil {
				serveError = responseErr
			}
			server.requestsServed += 1
			responseString := discoveryResponse.ToString() + "\n"
			fmt.Println("Response : ", responseString)
			tcpConnection.Write([]byte(responseString))
		}
	}
}
