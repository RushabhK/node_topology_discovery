package config_loader

import (
	"github.com/stretchr/testify/suite"
	"node_topology_discovery/model"
	"testing"
)

type ConfigDataTest struct {
	suite.Suite
}

func TestConfigDataTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigDataTest))
}

func (suite ConfigDataTest) TestShouldConvertConfigDataToNodeInfo() {
	configData := ConfigData{
		Port:      "3200",
		IpAddress: "localhost",
		Name:      "machine-0",
		Neighbors: []Neighbor{
			{
				IpAddress: "localhost",
				Port:      "3201",
			},
		},
	}

	expectedNodeInfo := model.NodeInfo{
		Port:      "3200",
		IpAddress: "localhost",
		Name:      "machine-0",
		Neighbors: []model.NeighborInfo{
			{
				IpAddress: "localhost",
				Port:      "3201",
			},
		},
	}

	suite.Equal(expectedNodeInfo, configData.ToNodeInfo())
}
