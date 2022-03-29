package main

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	authRouting "github.com/schiller-sql/littSQL/auth/delivery/routing"
	authR "github.com/schiller-sql/littSQL/auth/repository"
	authU "github.com/schiller-sql/littSQL/auth/usecase"
	"github.com/schiller-sql/littSQL/config"
	databasesRouting "github.com/schiller-sql/littSQL/databases/delivery/routing"
	databasesR "github.com/schiller-sql/littSQL/databases/repository"
	databasesU "github.com/schiller-sql/littSQL/databases/usecase"
	projectsRouting "github.com/schiller-sql/littSQL/projects/delivery/routing"
	projectsR "github.com/schiller-sql/littSQL/projects/repository"
	projectsU "github.com/schiller-sql/littSQL/projects/usecase"
	"github.com/spf13/viper"
)

func main() {
	// TODO: Set all routes here, so that auth middleware does not have to be given through
	config.InitConfigFile()

	mode := viper.Get("MODE").(string)
	gin.SetMode(mode)

	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	r.Use(config.InitCORSMiddleware())

	db := config.InitPostgresDB()

	authRepo := authR.NewRepository(db, viper.Get("BCRYPT_COST").(int))
	authUsecase := authU.NewUsecase(authRepo)
	authMiddleware := authM.NewAuthMiddleware(authUsecase)
	authRouting.ConfigureHandler(r, authMiddleware, authUsecase)

	databasesRepo := databasesR.NewRepository(db)
	databasesUsecase := databasesU.NewUsecase(databasesRepo)
	databasesRouting.ConfigureHandler(r, authMiddleware, databasesUsecase)

	projectsRepo := projectsR.NewRepository(db)
	projectsUsecase := projectsU.NewUsecase(projectsRepo, databasesRepo)
	projectsRouting.ConfigureHandler(r, authMiddleware, projectsUsecase)

	err = r.Run(":" + (viper.Get("PORT").(string)))
	if err != nil {
		panic(err)
	}
}
