// bukumigrate.go
package main

import (
    "log"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"buku-api-gin/config"
	"buku-api-gin/models"
)

func main() {
    _, err := config.InitDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer config.DB.Close()

    config.DB.AutoMigrate(&models.Bukus{})

    log.Println("Migrasi tabel buku selesai")
}
