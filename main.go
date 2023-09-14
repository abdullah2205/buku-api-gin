package main

import (
    "fmt"
    "log"
    "net/http"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    db  *gorm.DB
    err error
)

type bukus struct {
    ID     uint   `json:"id" gorm:"primary_key"`
    Judul  string `json:"judul"`
    Tahun string `json:"tahun"`
    UserID uint `json:"user_id"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
    UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

func main() {
    db, err = gorm.Open("mysql", "root:user@tcp(localhost:3306)/buku_api?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.AutoMigrate(&bukus{})

    router := gin.Default()
    router.GET("/api/buku", IndexBuku)
    router.POST("/api/buku", StoreBuku)
    router.GET("/api/buku/:id", ShowBuku)
    router.PUT("/api/buku/:id", UpdateBuku)
    router.DELETE("/api/buku/:id", DestroyBuku)
    router.Run(":8080")
}

func IndexBuku(c *gin.Context) {
    var books []bukus
    db.Find(&books)

    if len(books) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, books)
}

func StoreBuku(c *gin.Context) {
    var book bukus

    c.BindJSON(&book)
    db.Create(&book)

    c.JSON(http.StatusCreated, book)
}

func ShowBuku(c *gin.Context) {
    id := c.Param("id")
    var book bukus
    
    if err := db.Where("id = ?", id).First(&book).Error; err != nil {
        handleError(c, err)
        return
    }

    c.JSON(http.StatusOK, book)
}

func UpdateBuku(c *gin.Context) {
    id := c.Param("id")
    var book bukus

    if err := db.Where("id = ?", id).First(&book).Error; err != nil {
        handleError(c, err)
        return
    }

    c.BindJSON(&book)
    db.Save(&book)

    c.JSON(http.StatusOK, book)
}

func DestroyBuku(c *gin.Context) {
    id := c.Param("id")
    var book bukus

    if err := db.Where("id = ?", id).First(&book).Error; err != nil {
        handleError(c, err)
        return
    }

    db.Where("id = ?", id).Delete(&book)
    
    c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}

func handleError(c *gin.Context, err error) {
    fmt.Println("Kesalahan:", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
