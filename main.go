package main

import (
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	"buku-api-gin/config"
	"buku-api-gin/routes"
)

func main() {
	_, err := config.InitDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer config.DB.Close()

    router := gin.Default()
	routes.SetupRoutes(router)
    router.Run(":8080")
}



//tambah ini aja