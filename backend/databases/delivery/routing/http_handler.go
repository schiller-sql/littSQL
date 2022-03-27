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
	// TODO: Security flaw: only teacher or students of a course where the database is used should have access
	group.GET("/:id", jwtMiddleware.MiddlewareFunc(), handler.getDatabase)
	group.PUT("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.editDatabase)
	group.DELETE("/:id", jwtMiddleware.MiddlewareFunc(), teacherMiddleware, handler.deleteDatabase)
}

func getTeacherIDHelper(c *gin.Context) int32 {
	id, _ := c.Get("id")
	return id.(int32)
}

func getDatabaseIDHelper(c *gin.Context) (int32, error) {
	databaseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}
	return int32(databaseID), nil
}

func (h *databasesHandler) getDatabasesOfTeacher(c *gin.Context) {
	teacherID := getTeacherIDHelper(c)
	databasesOfTeacher, err := h.usecase.GetDatabasesOfTeacher(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, databasesOfTeacher)
}

func (h *databasesHandler) newDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databasesHandler) getDatabase(c *gin.Context) {
	databaseID, err := getDatabaseIDHelper(c)
	if err != nil {
		return
	}
	database, err := h.usecase.GetDatabaseDetails(databaseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	databaseForm := DatabaseForm{
		ID:            databaseID,
		Name:          database.Name,
		Data:          database.Data,
		SchemaSVGPath: database.SchemaSVGPath,
		IsPublic:      database.OwnerID == nil,
	}
	c.JSON(http.StatusOK, databaseForm)
}

func (h *databasesHandler) editDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databasesHandler) deleteDatabase(c *gin.Context) {
	panic("implement me")
}

type DatabaseForm struct {
	ID            int32   `json:"id"`
	Name          string  `json:"name"`
	Data          string  `json:"data"`
	SchemaSVGPath *string `json:"schema_svg_path"`
	IsPublic      bool    `json:"is_public"`
}
