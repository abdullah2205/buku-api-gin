package routes

import (
    "github.com/gin-gonic/gin"

    "buku-api-gin/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/register", controllers.Register)
    router.POST("/api/login", controllers.Login)

	//get user
	router.GET("/api/buku", controllers.AuthMiddleware(), controllers.IndexBuku)
	router.POST("/api/buku", controllers.StoreBuku)
	router.GET("/api/buku/:id", controllers.ShowBuku)
	router.PUT("/api/buku/:id", controllers.UpdateBuku)
	router.DELETE("/api/buku/:id", controllers.DestroyBuku)

	//logout
}