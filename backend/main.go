package main

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	authRouter "github.com/schiller-sql/littSQL/auth/delivery/routing"
	authR "github.com/schiller-sql/littSQL/auth/repository"
	authU "github.com/schiller-sql/littSQL/auth/usecase"
	"github.com/schiller-sql/littSQL/config"
	"github.com/spf13/viper"
)

func main() {
	// TODO: Set all routes here, so that auth middleware does not have to be given through
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	config.InitConfigFile()
	r.Use(config.InitCORSMiddleware())

	db := config.InitPostgresDB()

	authRepo := authR.NewRepository(db, viper.Get("BCRYPT_COST").(int))
	authUsecase := authU.NewUsecase(authRepo)
	authMiddleware := authM.NewAuthMiddleware(authUsecase)
	authRouter.ConfigureHandler(r, authMiddleware, authUsecase)

	err = r.Run(":" + (viper.Get("PORT").(string)))
	if err != nil {
		panic(err)
	}
}
