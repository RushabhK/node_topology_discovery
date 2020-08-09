package client

type ClientFactory interface {
	GetClient() Client
}
