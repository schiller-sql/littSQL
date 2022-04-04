package main

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	authRouting "github.com/schiller-sql/littSQL/auth/delivery/routing"
	authR "github.com/schiller-sql/littSQL/auth/repository"
	authU "github.com/schiller-sql/littSQL/auth/usecase"
	"github.com/schiller-sql/littSQL/config"
	databaseTemplatesRouting "github.com/schiller-sql/littSQL/database_templates/delivery/routing"
	databaseTemplatesR "github.com/schiller-sql/littSQL/database_templates/repository"
	databaseTemplatesU "github.com/schiller-sql/littSQL/database_templates/usecase"
	projectsRouting "github.com/schiller-sql/littSQL/projects/delivery/routing"
	projectsR "github.com/schiller-sql/littSQL/projects/repository"
	projectsU "github.com/schiller-sql/littSQL/projects/usecase"
	sqlExecutorR "github.com/schiller-sql/littSQL/sql_executor/repository"
	"github.com/spf13/viper"
)

func getRouter() *gin.Engine {
	mode := viper.Get("MODE").(string)
	gin.SetMode(mode)

	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	r.Use(config.InitCORSMiddleware())
	return r
}

func main() {
	// TODO: Set all routes here, so that auth middleware does not have to be given through
	config.InitConfigFile()

	r := getRouter()

	db := config.InitPostgresDB()

	authRepo := authR.NewRepository(db, viper.Get("BCRYPT_COST").(int))
	databaseTemplatesRepo := databaseTemplatesR.NewRepository(db)
	projectsRepo := projectsR.NewRepository(db)
	sqlExecutorRepo := sqlExecutorR.NewRepository(viper.Get("SQLITE_DIR").(string))

	authUsecase := authU.NewUsecase(authRepo)
	authMiddleware := authM.NewAuthMiddleware(authUsecase)
	authRouting.ConfigureHandler(r, authMiddleware, authUsecase)

	databaseTemplatesUsecase := databaseTemplatesU.NewUsecase(databaseTemplatesRepo)
	databaseTemplatesRouting.ConfigureHandler(r, authMiddleware, databaseTemplatesUsecase)

	projectsUsecase := projectsU.NewUsecase(projectsRepo, sqlExecutorRepo)
	projectsRouting.ConfigureHandler(r, authMiddleware, projectsUsecase)

	if err := r.Run(":" + (viper.Get("PORT").(string))); err != nil {
		panic(err)
	}
}
