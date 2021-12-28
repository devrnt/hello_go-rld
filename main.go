package main

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/routes"
)

func main() {
	var router = gin.Default()

	routes.InstallAlbumRoutes(router)

	router.Run("localhost:8080")
}
