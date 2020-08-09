package client

import "time"

type udpClientFactory struct {
	timeoutInSeconds time.Duration
}

func NewUdpClientFactory(timeoutInSeconds time.Duration) ClientFactory {
	return udpClientFactory{
		timeoutInSeconds: timeoutInSeconds,
	}
}

func (factory udpClientFactory) GetClient() Client {
	return udpClient{
		timeoutInSeconds: factory.timeoutInSeconds,
	}
}
