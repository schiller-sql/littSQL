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
	"github.com/spf13/viper"
	"net/http"
)

func getRouter() *gin.Engine {
	mode := viper.Get("MODE").(string)
	gin.SetMode(mode)

	router := gin.Default()

	// setup static files for svelte
	router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("../frontend/public", false))))

	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	return router
}

func main() {
	config.InitConfigFile()
	db := config.InitPostgresDB()

	r := getRouter()
	{
		g := r.Group("/api")

		authRepo := authR.NewRepository(db, viper.Get("BCRYPT_COST").(int))
		authUsecase := authU.NewUsecase(authRepo)
		authMiddleware := authM.NewAuthMiddleware(authUsecase)
		authRouting.ConfigureHandler(g, authMiddleware, authUsecase)

		databaseTemplatesRepo := databaseTemplatesR.NewRepository(db)
		databaseTemplatesUsecase := databaseTemplatesU.NewUsecase(databaseTemplatesRepo)
		databaseTemplatesRouting.ConfigureHandler(g, authMiddleware, databaseTemplatesUsecase)

		projectsRepo := projectsR.NewRepository(db)
		projectsUsecase := projectsU.NewUsecase(projectsRepo)
		projectsRouting.ConfigureHandler(g, authMiddleware, projectsUsecase)
	}

	err := r.Run(":" + (viper.Get("PORT").(string)))
	if err != nil {
		panic(err)
	}
}
