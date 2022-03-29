package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() *gorm.DB {
	host := viper.Get("PG_HOST")
	name := viper.Get("PG_NAME")
	user := viper.Get("PG_USER")
	pass := viper.Get("PG_PASSWORD")

	dsn := fmt.Sprintf("host=%v database=%v user=%v password=%v", host, name, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
