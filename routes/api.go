package routes

import (
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"


    "buku-api-gin/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Membuat instance middleware CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://192.168.1.15:8080"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}

	// Menggunakan middleware Cors
	router.Use(cors.New(corsConfig))

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