package controller

import (
	"demo_hubu_backend/middleware"
	"demo_hubu_backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context, db *gorm.DB) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser model.User
	log.Printf("user is %+v", user)
	if err := middleware.DB().Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		log.Printf("Database query error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	log.Printf("existingUser is %+v", existingUser)
	if existingUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := generateToken(existingUser.ID, existingUser.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register 用户注册
func Register(c *gin.Context, db *gorm.DB) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := middleware.DB().Create(&user).Error; err != nil {
		log.Printf("Database insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Logout(c *gin.Context, db *gorm.DB) {
	fmt.Println("logout")
}

// generateToken 生成 JWT Token
func generateToken(userID uint, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(middleware.JWTSecret()))
	return tokenString
}
