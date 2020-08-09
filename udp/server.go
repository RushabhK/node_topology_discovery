package udp

import (
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/model"
	"node_topology_discovery/service"
	"strconv"
	"sync"
)

type Server interface {
	Serve(port string, nodesCount chan int, wg *sync.WaitGroup) error
}

type udpServer struct {
	requestsServed  int
	totalNodesCount int
	service         service.NodesDiscoveryService
}

func NewUdpServer(discoveryService service.NodesDiscoveryService) Server {
	return udpServer{
		service:         discoveryService,
		requestsServed:  0,
		totalNodesCount: 0,
	}
}

func (server udpServer) Serve(port string, nodesCount chan int, wg *sync.WaitGroup) (serveError error) {
	defer wg.Done()

	portInt, _ := strconv.Atoi(port)
	addr := net.UDPAddr{
		Port: portInt,
		IP:   net.ParseIP("localhost"),
	}
	udpServer, listenError := net.ListenUDP("udp", &addr)
	if listenError != nil {
		fmt.Println("Error while listening : ", listenError.Error())
		serveError = listenError
	}
	requestBytes := make([]byte, 2048)

	for {
		if len(nodesCount) > 0 {
			server.totalNodesCount = <-nodesCount
		}
		if server.totalNodesCount != 0 && server.requestsServed == server.totalNodesCount-1 {
			return nil
		}

		println("Reading request in server")
		readLen, remoteaddr, err := udpServer.ReadFromUDP(requestBytes)
		if err != nil {
			return err
		}
		requestString := string(requestBytes)
		println("Request from client: ", requestString)
		println("Remote address: ", remoteaddr)

		var request model.NodesDiscoveryRequest
		unmarshalErr := json.Unmarshal(requestBytes[0:readLen], &request)
		if unmarshalErr != nil {
			serveError = unmarshalErr
			println("Unmarshalling error : ", unmarshalErr.Error())
			return
		}


		fmt.Println("Request received : ", request.ToString())

		if len(request.VisitedNodes) > 0 {
			discoveryResponse, responseErr := server.service.Discover(request)
			if responseErr != nil {
				serveError = responseErr
				return
			}
			server.requestsServed += 1
			responseString := discoveryResponse.ToString() + "\n"
			fmt.Println("Response : ", responseString)
			udpServer.WriteToUDP([]byte(responseString), remoteaddr)
		}
	}
}
