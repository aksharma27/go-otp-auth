
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aksharma27/auth/internal/config"
)

func main() {
	// _ = godotenv.Load()

	cfg := config.LoadConfig()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "OTP Auth Service is running",
		})
	})

	router.Run(":" + cfg.AppPort)
}
