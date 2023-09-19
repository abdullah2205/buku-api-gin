package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	"buku-api-gin/config"
	"buku-api-gin/models"
)

func IndexBuku(c *gin.Context) {
    var buku []models.Bukus

    if err := config.DB.Find(&buku).Error; err != nil {
        errorMsg := err.Error()

        c.JSON(http.StatusInternalServerError, gin.H{"error": errorMsg})
        return
    }

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
    var buku models.Bukus

    c.BindJSON(&buku)
    config.DB.Create(&buku)

    tambah_buku := gin.H{
        "_pesan": "Buku berhasil ditambah",
        "data": buku,
    }

    c.JSON(http.StatusCreated, tambah_buku)
}