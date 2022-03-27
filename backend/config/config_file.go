package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfigFile() {
	viper.SetDefault("MODE", "debug")

	viper.SetDefault("PORT", "8080")
	// TODO: add underscores to make
	viper.SetDefault("PGHOST", "127.0.0.1")
	viper.SetDefault("PGNAME", "postgres")
	viper.SetDefault("PGUSER", "")
	viper.SetDefault("PGPASSWORD", "postgres")

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
