package config_loader

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigDataTestSuite struct {
	suite.Suite
	configData ConfigData
}

func TestConfigDataTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigDataTestSuite))
}

func (suite *ConfigDataTestSuite) SetupTest() {
	//suite.configData = ConfigData{
	//	Profiles: []ProfileConfig{
	//		{
	//			Name:               "work",
	//			BookmarksDirectory: "work-dir",
	//		},
	//		{
	//			Name:               "test",
	//			BookmarksDirectory: "test-dir",
	//		},
	//	},
	//}
}

func (suite ConfigDataTestSuite) TestShouldReturnProfileGivenTheName() {
	//profile, profileError := suite.configData.GetProfile("work")
	//
	//suite.Nil(profileError)
	//suite.Equal(ProfileConfig{
	//	Name:               "work",
	//	BookmarksDirectory: "work-dir",
	//}, profile)
}

func (suite ConfigDataTestSuite) TestShouldReturnProfileNotPresentErrorWhenProfileIsNotPresent() {
	//profile, profileError := suite.configData.GetProfile("random-profile")
	//
	//suite.Equal(server_errors.ProfileNotPresent{}, profileError)
	//suite.Empty(profile)
}
