package routing

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/databases"
	"net/http"
	"strconv"
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
	group.PUT("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.editDatabase)
	group.DELETE("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.deleteDatabase)
}

func getDatabaseIDHelper(c *gin.Context) (int32, error) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}
	return int32(projectID), nil
}

func (h *databasesHandler) getDatabasesOfTeacher(c *gin.Context) {
	panic("implement me")
}

func (h *databasesHandler) newDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databasesHandler) getDatabase(c *gin.Context) {
	id, err := getDatabaseIDHelper(c)
	if err != nil {
		return
	}
	database, err := h.usecase.GetDatabaseDetails(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, database)
}

func (h *databasesHandler) editDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databasesHandler) deleteDatabase(c *gin.Context) {
	panic("implement me")
}
