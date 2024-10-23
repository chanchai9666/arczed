package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"arczed/pkg/safety"
)

// AuthMiddleware เป็น middleware สำหรับตรวจสอบ JWT และ Basic Auth
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ตรวจสอบ Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Please Login before use."})
			c.Abort()
			return
		}

		// ตรวจสอบประเภทการเข้าสู่ระบบ
		if strings.HasPrefix(authHeader, "Bearer ") {
			// JWT Authentication
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := safety.VerifyJWT(jwtSecret, tokenString) // เรียกใช้ฟังก์ชัน VerifyJWT

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
				c.Abort()
				return
			}

			// เก็บ claims ใน context เพื่อใช้งานต่อไป
			c.Set("claims", claims)
			c.Set("userID", claims.UserId)

		} else if strings.HasPrefix(authHeader, "Basic ") {
			// Basic Authentication
			encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
			decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Basic Auth"})
				c.Abort()
				return
			}

			// แยก username และ password
			credentials := strings.SplitN(string(decodedCredentials), ":", 2)
			if len(credentials) != 2 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Basic Auth"})
				c.Abort()
				return
			}

			username, password := credentials[0], credentials[1]

			// ตรวจสอบ username และ password (แทนที่ด้วยการตรวจสอบจริง)
			if !validateBasicAuth(username, password) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unsupported Authorization method"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// validateBasicAuth เป็นฟังก์ชันตัวอย่างสำหรับตรวจสอบ username และ password
func validateBasicAuth(username, password string) bool {
	// แทนที่ด้วยการตรวจสอบจริง (เช่นจากฐานข้อมูล)
	return username == "admin" && password == "admin"
}

func GetUserProfile(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		return ""
	}

	return userID.(string) // แปลงค่าให้เป็น string
}
