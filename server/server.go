package server

import (
	"sync"
)

type Server interface {
	Serve(port string, nodesCount chan int, wg *sync.WaitGroup) error
}
