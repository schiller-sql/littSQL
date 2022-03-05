package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfigFile() {
	viper.SetDefault("PGHOST", "127.0.0.1")
	viper.SetDefault("PGNAME", "postgres")
	viper.SetDefault("PGUSER", "")
	viper.SetDefault("PGPASSWORD", "postgres")

	viper.SetDefault("CORS_ORIGIN", "http://127.0.0.1")

	viper.SetDefault("JWT_SECRET", "1")
	viper.SetDefault("JWT_SIGN_ALG", "HS256")

	viper.SetDefault("BCRYPT_COST", 4)

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
