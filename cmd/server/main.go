
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aksharma27/go-otp-auth/internal/config"
	"github.com/aksharma27/go-otp-auth/internal/database"
	"github.com/aksharma27/go-otp-auth/internal/repository"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	_ = userRepo

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			2*time.Second,
		)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"success": false,
				"message": "Database unavailable",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Service and database are healthy",
		})
	})

	log.Fatal(router.Run(":" + cfg.AppPort))
}
