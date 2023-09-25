package controllers

import (
    "fmt"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"

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

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"kesalahan": err.Error()})
        return
    }

    token, err := GenerateJWTToken(fmt.Sprint(user.ID)) //kirim ke generate

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    login := gin.H{
        "_pesan": "Berhasil Login",
        "data": user,
        "token" : token,
    }

    c.JSON(http.StatusCreated, login)
}

//Authorization logic
var (
	secretKey = []byte("ini_kunci")
)

func GenerateJWTToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not Found"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if !token.Valid {
			fmt.Print("error : ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

        claims, _ := token.Claims.(jwt.MapClaims)
        c.Set("user_id", claims["user_id"])

		c.Next()
	}
}