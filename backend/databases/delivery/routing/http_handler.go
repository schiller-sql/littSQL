package routing

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/helpers"
	"net/http"
)

type databasesHandler struct {
	usecase databases.Usecase
}

func ConfigureHandler(r *gin.Engine, authMiddleware *authM.AuthMiddleware, usecase databases.Usecase) {
	handler := databasesHandler{usecase}

	group := r.Group("/databases", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	// TODO: implement other methods, or let it be
	group.GET("/", handler.getDatabasesOfTeacher)
	//group.POST("/", handler.newDatabase)
	group.GET("/:id", handler.getDatabase)
	//group.PUT("/:id", handler.editDatabase)
	//group.DELETE("/:id", handler.deleteDatabase)
}

func (h *databasesHandler) getDatabasesOfTeacher(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
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
	databaseID, err := helpers.GetParamID(c)
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
