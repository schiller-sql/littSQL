package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitCORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Authorization"}
	config.AllowOrigins = []string{viper.Get("CORS_ORIGIN").(string)}
	return cors.New(config)
}
