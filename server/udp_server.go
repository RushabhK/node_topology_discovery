package server

import (
	"encoding/json"
	"fmt"
	"net"
	"node_topology_discovery/model"
	"node_topology_discovery/service"
	"strconv"
	"sync"
)

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

type udpPayload struct {
	data          []byte
	remoteAddress *net.UDPAddr
}

func readFromUdp(conn *net.UDPConn, readChan chan udpPayload) error {
	requestBytes := make([]byte, 2048)
	for {
		fmt.Println("Reading request from the client")
		if conn == nil {
			return nil
		}
		readLen, remoteaddr, err := conn.ReadFromUDP(requestBytes)
		if err != nil {
			fmt.Println("Error while reading the request : ", err.Error())
			return err
		}
		fmt.Println("Request from client: ", string(requestBytes))
		fmt.Println("Remote address: ", remoteaddr)
		readChan <- udpPayload{data: requestBytes[0:readLen], remoteAddress: remoteaddr}
	}
}

func (server udpServer) Serve(port string, nodesCountChan chan int, wg *sync.WaitGroup) (serveError error) {
	defer wg.Done()

	portInt, _ := strconv.Atoi(port)
	addr := net.UDPAddr{
		Port: portInt,
		IP:   net.ParseIP("localhost"),
	}
	fmt.Println("Listening to udp..")
	udpServer, listenError := net.ListenUDP("udp", &addr)
	if listenError != nil {
		fmt.Println("Error while listening : ", listenError.Error())
		serveError = listenError
	}
	udpPayloadChan := make(chan udpPayload, 1000)
	go readFromUdp(udpServer, udpPayloadChan)

	for {
		fmt.Println("Total nodes count detected by server: ", server.totalNodesCount)
		fmt.Println("Total requests served by server: ", server.requestsServed)
		if server.totalNodesCount != 0 && server.requestsServed >= server.totalNodesCount-1 {
			fmt.Println("Processing of the server is completed.")
			return nil
		}

		select {
		case nodesCount := <-nodesCountChan:
			{
				fmt.Println("Length of Channel: ", len(nodesCountChan))
				server.totalNodesCount = nodesCount
				fmt.Println("Total nodes in the topology : ", server.totalNodesCount)
			}
		case udpPayload := <-udpPayloadChan:
			{
				fmt.Println("Unmarshalling the request..")
				var request model.NodesDiscoveryRequest
				unmarshalErr := json.Unmarshal(udpPayload.data, &request)
				if unmarshalErr != nil {
					println("Unmarshalling error : ", unmarshalErr.Error())
					serveError = unmarshalErr
					return
				}

				fmt.Println("Request unmarshalling successful : ", request.ToString())

				if len(request.VisitedNodes) > 0 {
					fmt.Println("Valid request")
					discoveryResponse, responseErr := server.service.Discover(request)
					if responseErr != nil {
						fmt.Println("Error while discovering ", responseErr.Error())
						serveError = responseErr
						return
					}
					server.requestsServed += 1
					responseString := discoveryResponse.ToString() + "\n"
					fmt.Println("Response from discovery : ", responseString)
					fmt.Println("Writing the discovery response..")
					_, writeError := udpServer.WriteToUDP([]byte(responseString), udpPayload.remoteAddress)
					if writeError != nil {
						fmt.Println("Error while writing the response from the server : ", writeError.Error())
					}
				}
			}
		}
	}

}
