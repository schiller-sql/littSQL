package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfigFile() {
	viper.SetDefault("MODE", "debug")

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("PG_HOST", "127.0.0.1")
	viper.SetDefault("PG_NAME", "postgres")
	viper.SetDefault("PG_USER", "")
	viper.SetDefault("PG_PASSWORD", "postgres")

	viper.SetDefault("CORS_ORIGIN", "http://localhost:5555")

	viper.SetDefault("JWT_SECRET", "1")
	viper.SetDefault("JWT_SIGN_ALG", "HS256")

	viper.SetDefault("BCRYPT_COST", 4)

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(".env file not found, using defaults...")
	}
}
