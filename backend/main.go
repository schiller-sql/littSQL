package main

import (
	"github.com/gin-gonic/gin"
	authH "github.com/schiller-sql/littSQL/auth/delivery/http_handler"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	authR "github.com/schiller-sql/littSQL/auth/repository"
	authU "github.com/schiller-sql/littSQL/auth/usecase"
	"github.com/schiller-sql/littSQL/config"
	databasesH "github.com/schiller-sql/littSQL/databases/delivery/http_handler"
	databasesR "github.com/schiller-sql/littSQL/databases/repository"
	databasesU "github.com/schiller-sql/littSQL/databases/usecase"
	projectsH "github.com/schiller-sql/littSQL/projects/delivery/http_handler"
	projectsR "github.com/schiller-sql/littSQL/projects/repository"
	projectsU "github.com/schiller-sql/littSQL/projects/usecase"
	"github.com/spf13/viper"
)

func getRouterRef() *gin.Engine {
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

func setUpAuth(group *gin.RouterGroup, handler authH.HttpHandler, authMiddleware *authM.AuthMiddleware) {
	group.POST("/signup", handler.Signup)
	group.POST("/login", authMiddleware.LoginHandler)
	group.POST("/logout", authMiddleware.LogoutHandler)
	group.GET("/refresh_token", authMiddleware.RefreshHandler)
	group.GET("/account", authMiddleware.JwtHandler, handler.GetAccountDetails)
	group.DELETE("/account", authMiddleware.JwtHandler, handler.DeleteAccount)
}

func setUpDatabases(group *gin.RouterGroup, handler databasesH.HttpHandler, authMiddleware *authM.AuthMiddleware) {
	group.GET("/", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator, handler.GetDatabasesOfTeacher)
	group.POST("/", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator, handler.NewDatabase)
	// TODO: Security flaw: only teacher or students of a course where the database is used should have access
	group.GET("/:id", authMiddleware.JwtHandler, handler.GetDatabase)
	group.PUT("/:id", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator, handler.EditDatabase)
	group.DELETE("/:id", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator, handler.DeleteDatabase)
}

func setUpProjects(group *gin.RouterGroup, handler projectsH.HttpHandler, authMiddleware *authM.AuthMiddleware) {
	group.GET("/", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator, handler.GetProjectsOfTeacher)
	group.POST("/", authMiddleware.IsTeacherValidator, handler.NewProject)
	group.GET("/:id", authMiddleware.IsTeacherValidator, handler.GetProject)
	group.PUT("/:id", authMiddleware.IsTeacherValidator, handler.EditProject)
	group.DELETE("/:id", authMiddleware.IsTeacherValidator, handler.DeleteProject)
}

func main() {
	// TODO: Set all routes here, so that auth middleware does not have to be given through
	config.InitConfigFile()

	r := getRouterRef()

	db := config.InitPostgresDB()

	// repositories
	authRepo := authR.NewRepository(db, viper.Get("BCRYPT_COST").(int))
	databasesRepo := databasesR.NewRepository(db)
	projectsRepo := projectsR.NewRepository(db)

	// usecases
	authUsecase := authU.NewUsecase(authRepo)
	databasesUsecase := databasesU.NewUsecase(databasesRepo)
	projectsUsecase := projectsU.NewUsecase(projectsRepo, databasesRepo)

	// middleware
	authMiddleware := authM.NewAuthMiddleware(authUsecase)

	// http routing
	authHandler := authH.NewHttpHandler(authUsecase)
	authGroup := r.Group("/auth")
	setUpAuth(authGroup, authHandler, authMiddleware)

	databasesHandler := databasesH.NewHttpHandler(databasesUsecase)
	databasesGroup := r.Group("/databases")
	setUpDatabases(databasesGroup, databasesHandler, authMiddleware)

	projectsHandler := projectsH.NewHttpHandler(projectsUsecase)
	projectsGroup := r.Group("/projects")
	setUpProjects(projectsGroup, projectsHandler, authMiddleware)

	err := r.Run(":" + (viper.Get("PORT").(string)))
	if err != nil {
		panic(err)
	}
}
