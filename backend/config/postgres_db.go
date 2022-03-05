package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() *gorm.DB {
	host := viper.Get("PGHOST")
	name := viper.Get("PGNAME")
	user := viper.Get("PGUSER")
	pass := viper.Get("PGPASSWORD")

	dsn := fmt.Sprintf("host=%v database=%v user=%v password=%v", host, name, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
