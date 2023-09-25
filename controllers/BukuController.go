package controllers

import (
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"buku-api-gin/config"
	"buku-api-gin/models"
)

func IndexBuku(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var buku []models.Bukus

    if err := config.DB.Where("user_id = ?", userID).Find(&buku).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"kesalahan": err.Error()})
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
    userID, _ := c.Get("user_id")

    var buku models.Bukus

    userIDUint, _ := strconv.ParseUint(userID.(string), 10, 64)

    buku.UserID = userIDUint

    if err := c.BindJSON(&buku); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"kesalahan": err.Error()})
        return
    }
    
    if err := config.DB.Create(&buku).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"kesalahan": err.Error()})
        return
    }

    tambah_buku := gin.H{
        "_pesan": "Buku berhasil ditambah",
        "data": buku,
    }

    c.JSON(http.StatusCreated, tambah_buku)
}

func ShowBuku(c *gin.Context) {
    userID, _ := c.Get("user_id")

    id := c.Param("id")
    var buku models.Bukus
    
    if err := config.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"kesalahan": err.Error()})
        return
    }

    data_buku := gin.H{
        "_pesan": "Data Buku",
        "data": buku,
    }

    c.JSON(http.StatusOK, data_buku)
}

func UpdateBuku(c *gin.Context) {
    userID, _ := c.Get("user_id")

    id := c.Param("id")
    var buku models.Bukus

    if err := config.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"kesalahan": err.Error()})
        return
    }

    if err := c.BindJSON(&buku); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"kesalahan": err.Error()})
        return
    }

    if err := config.DB.Save(&buku).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"kesalahan": err.Error()})
        return
    }

    ubah_buku := gin.H{
        "_pesan": "Buku berhasil diubah",
        "data": buku,
    }

    c.JSON(http.StatusOK, ubah_buku)
}

func DestroyBuku(c *gin.Context) {
    userID, _ := c.Get("user_id")

    id := c.Param("id")
    var buku models.Bukus

    if err := config.DB.Where("id = ?", id).Where("user_id = ?", userID).First(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"kesalahan": err.Error()})
        return
    }

    if err := config.DB.Where("id = ?", id).Delete(&buku).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"kesalahan": err.Error()})
        return   
    }

    hapus_buku := gin.H{
        "_pesan": "Buku berhasil dihapus",
        "data": buku,
    }
    
    c.JSON(http.StatusOK, hapus_buku)
}