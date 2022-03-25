package routing

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/databases"
	"net/http"
)

type databasesHandler struct {
	usecase databases.Usecase
}

func ConfigureHandler(r *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware, usecase databases.Usecase) {
	handler := databasesHandler{usecase}

	group := r.Group("/databases")

	teacherMiddleware := func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if !claims["is_teacher"].(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You have to be a teacher to access this resource"})
		}
	}

	group.GET("/", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.getDatabasesOfTeacher)
	group.POST("/", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.newDatabase)
	group.GET("/:id", jwtMiddleware.MiddlewareFunc(), handler.getDatabase)
	group.POST("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.editDatabase)
	group.DELETE("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.deleteDatabase)
}

func (h *databasesHandler) getDatabasesOfTeacher(c *gin.Context) {}
func (h *databasesHandler) newDatabase(c *gin.Context)           {}
func (h *databasesHandler) getDatabase(c *gin.Context)           {}
func (h *databasesHandler) editDatabase(c *gin.Context)          {}
func (h *databasesHandler) deleteDatabase(c *gin.Context)        {}
