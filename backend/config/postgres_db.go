package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

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
	var logLevel logger.LogLevel
	if viper.Get("mode") == "debug" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}
