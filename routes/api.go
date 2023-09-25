package routes

import (
    "github.com/gin-gonic/gin"

    "buku-api-gin/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/api/register", controllers.Register)
    router.POST("/api/login", controllers.Login)

	auth := router.Group("/api")
	auth.Use(controllers.AuthMiddleware())

	auth.GET("/buku", controllers.IndexBuku)
	auth.POST("/buku", controllers.StoreBuku)
	auth.GET("/buku/:id", controllers.ShowBuku)
	auth.PUT("/buku/:id", controllers.UpdateBuku)
	auth.DELETE("/buku/:id", controllers.DestroyBuku)

	//logout
}