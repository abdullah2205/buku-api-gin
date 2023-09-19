package routes

import (
    "github.com/gin-gonic/gin"

    "buku-api-gin/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/buku", controllers.IndexBuku)
	router.POST("/api/buku", controllers.StoreBuku)
	// router.GET("/api/buku/:id", ShowBuku)
	// router.PUT("/api/buku/:id", UpdateBuku)
	// router.DELETE("/api/buku/:id", DestroyBuku)
}