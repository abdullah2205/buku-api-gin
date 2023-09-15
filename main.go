package main

import (
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
    var buku []bukus
    db.Find(&buku)

    if len(buku) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"pesan": "Buku tidak ditemukan"})
        return
    }

    list_buku := gin.H{
        "_pesan": "List Buku",
        "data": buku,
    }

    c.JSON(http.StatusOK, list_buku)
}

func StoreBuku(c *gin.Context) {
    var buku bukus

    c.BindJSON(&buku)
    db.Create(&buku)

    tambah_buku := gin.H{
        "_pesan": "Buku berhasil ditambah",
        "data": buku,
    }

    c.JSON(http.StatusCreated, tambah_buku)
}

func ShowBuku(c *gin.Context) {
    id := c.Param("id")
    var buku bukus
    
    if err := db.Where("id = ?", id).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"pesan": "Buku tidak ditemukan"})
        return
    }

    data_buku := gin.H{
        "_pesan": "Data Buku",
        "data": buku,
    }

    c.JSON(http.StatusOK, data_buku)
}

func UpdateBuku(c *gin.Context) {
    id := c.Param("id")
    var buku bukus

    if err := db.Where("id = ?", id).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"pesan": "Buku tidak ditemukan"})
        return
    }

    c.BindJSON(&buku)
    db.Save(&buku)
    //berikan validator nanti nya
    ubah_buku := gin.H{
        "_pesan": "Buku berhasil diubah",
        "data": buku,
    }

    c.JSON(http.StatusOK, ubah_buku)
}

func DestroyBuku(c *gin.Context) {
    id := c.Param("id")
    var buku bukus

    if err := db.Where("id = ?", id).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"pesan": "Buku tidak ditemukan"})
        return
    }

    db.Where("id = ?", id).Delete(&buku)

    hapus_buku := gin.H{
        "_pesan": "Buku berhasil dihapus cik",
        "data": buku,
    }
    
    c.JSON(http.StatusOK, hapus_buku)
}
