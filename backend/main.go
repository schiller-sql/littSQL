package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/config"
	uMiddle "github.com/schiller-sql/littSQL/users/delivery/middleware"
	uRout "github.com/schiller-sql/littSQL/users/delivery/routing"
	uRepos "github.com/schiller-sql/littSQL/users/repository"
	uUsec "github.com/schiller-sql/littSQL/users/usecase"
	"github.com/spf13/viper"
)

func main() {

	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	config.InitConfigFile()
	router.Use(config.InitCORSMiddleware())

	db := config.InitPostgresDB()

	usersRepository := uRepos.NewRepository(db, viper.Get("BCRYPT_COST").(int))
	usersUsecase := uUsec.NewUsecase(usersRepository)
	usersMiddleware := uMiddle.NewUsersMiddleware(usersUsecase)
	uRout.ConfigureHandler(router, usersMiddleware, usersUsecase)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
