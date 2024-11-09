package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB
var jwtSecret string

// SetJWTSecret 设置 JWT 密钥
func SetJWTSecret(secret string) {
	jwtSecret = secret
}

// JWTSecret 获取 JWT 密钥
func JWTSecret() string {
	return jwtSecret
}

// DB 获取数据库连接
func DB() *gorm.DB {
	if db == nil {
		log.Fatalf("Database connection is not set")
	}
	return db
}

// SetDB 设置数据库连接
func SetDB(d *gorm.DB) {
	db = d
}

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)
		c.Set("user_id", userID)
		c.Set("role", role)
		log.Printf("Authenticated user with ID: %d, Role: %s", userID, role)
		c.Next()
	}
}

// RoleMiddleware 角色中间件
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)
		log.Printf("Required role: %s, User role: %s", requiredRole, role)
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
