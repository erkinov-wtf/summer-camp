package main

import (
	"camp-summer/internal/api/dependencyIndejection"
	"camp-summer/internal/config"
	"camp-summer/internal/initializers"
	"camp-summer/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	config.MustLoad()

	initializers.ConnectDB()
	initializers.LoadTimezone()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(gin.Recovery())

	apiPrefixRoute := router.Group("/api/camp")

	container := dependencyIndejection.NewContainer()

	// Set up routes
	route.SetupAppRoutes(apiPrefixRoute, container)

	err := router.Run()
	if err != nil {
		return
	}
}
