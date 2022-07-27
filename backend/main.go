package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	authRouting "github.com/schiller-sql/littSQL/auth/delivery/routing"
	authR "github.com/schiller-sql/littSQL/auth/repository"
	authU "github.com/schiller-sql/littSQL/auth/usecase"
	"github.com/schiller-sql/littSQL/config"
	coursesRouting "github.com/schiller-sql/littSQL/courses/delivery/routing"
	coursesR "github.com/schiller-sql/littSQL/courses/repository"
	coursesU "github.com/schiller-sql/littSQL/courses/usecase"
	databaseTemplatesRouting "github.com/schiller-sql/littSQL/database_templates/delivery/routing"
	databaseTemplatesR "github.com/schiller-sql/littSQL/database_templates/repository"
	databaseTemplatesU "github.com/schiller-sql/littSQL/database_templates/usecase"
	participantsRouting "github.com/schiller-sql/littSQL/participants/delivery/routing"
	participantsR "github.com/schiller-sql/littSQL/participants/repository"
	participantsU "github.com/schiller-sql/littSQL/participants/usecase"
	projectsRouting "github.com/schiller-sql/littSQL/projects/delivery/routing"
	projectsR "github.com/schiller-sql/littSQL/projects/repository"
	projectsU "github.com/schiller-sql/littSQL/projects/usecase"
	"github.com/spf13/viper"
)

func getRouter() *gin.Engine {
	mode := viper.Get("MODE").(string)
	gin.SetMode(mode)

	router := gin.Default()

	// setup static files for svelte
	// TODO: use gzip
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

		coursesRepo := coursesR.NewRepository(db)
		coursesUsecase := coursesU.NewUsecase(coursesRepo)
		courseGroup := coursesRouting.ConfigureHandler(g, authMiddleware, coursesUsecase)

		participantsRepo := participantsR.NewRepository(db)
		participantsUsecase := participantsU.NewUsecase(participantsRepo, coursesRepo)
		participantsRouting.ConfigureHandler(courseGroup, authMiddleware, participantsUsecase)
	}

	err := r.Run(":" + (viper.Get("PORT").(string)))
	if err != nil {
		panic(err)
	}
}
