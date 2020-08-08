package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"node_topology_discovery/config_loader"
	"node_topology_discovery/constants"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "server is live..")
	})

	configLoader := config_loader.NewConfigLoader()
	configData, loadError := configLoader.Load(constants.CONFIG_FILE_PATH)

	if loadError != nil {
		panic("Cannot load config: " + loadError.Error())
	}

	r.Run(":" + configData.Port)
}
