package config_loader

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigLoaderTestSuite struct {
	suite.Suite
	configLoader ConfigLoader
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigLoaderTestSuite))
}

func (suite *ConfigLoaderTestSuite) SetupTest() {
	suite.configLoader = NewConfigLoader()
}

func (suite ConfigLoaderTestSuite) TestShouldLoadConfigFromFile() {
	configData, err := suite.configLoader.Load("./test_resources/config.json")

	suite.Nil(err)

	expectedConfig := ConfigData{
		Name:      "machine-1",
		IpAddress: "localhost",
		Port:      "3201",
		Neighbors: []Neighbor{
			{
				Name:      "machine-2",
				IpAddress: "localhost",
				Port:      "3202",
			},
			{
				Name:      "machine-3",
				IpAddress: "localhost",
				Port:      "3203",
			},
		},
	}
	suite.Equal(expectedConfig, configData)
}

func (suite ConfigLoaderTestSuite) TestShouldThrowErrorWhenFileNotPresent() {
	configData, err := suite.configLoader.Load("file-path-not-present.json")

	suite.NotNil(err)
	suite.Empty(configData)
}
