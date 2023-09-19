package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

	"buku-api-gin/config"
	"buku-api-gin/models"
)

func Register(c *gin.Context) {
	type InputValidator struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input InputValidator

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"kesalahan": err.Error()})
        return
    }

    user := models.User{
		Name: input.Name,
        Email: input.Email,
        Password: string(hashedPassword),
    }

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"kesalahan": err.Error()})
        return
    }

	register := gin.H{
        "_pesan": "Berhasil Register",
        "data": user,
    }

    c.JSON(http.StatusCreated, register)
}

func Login(c *gin.Context) {
    type InputValidator struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

    var input InputValidator

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"kesalahan": err.Error()})
        return
    }

    var user models.User
    
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"kesalahan": err.Error()})
        return
    }

    // Periksa password dengan hash yang ada di database
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"kesalahan": err.Error()})
        return
    }

    login := gin.H{
        "_pesan": "Berhasil Login",
        "data": user,
    }

    c.JSON(http.StatusCreated, login)
}
