package config_loader

type ConfigData struct {
	Name      string     `json:"name"`
	IpAddress string     `json:"ip_address"`
	Port      string     `json:"port"`
	Neighbors []Neighbor `json:"neighbors"`
}

type Neighbor struct {
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Port      string `json:"port"`
}
