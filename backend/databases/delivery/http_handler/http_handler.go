package http_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/helpers"
	"net/http"
)

type HttpHandler struct {
	usecase databases.Usecase
}

func NewHttpHandler(usecase databases.Usecase) HttpHandler {
	return HttpHandler{usecase}
}

func (h *HttpHandler) GetDatabasesOfTeacher(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	databasesOfTeacher, err := h.usecase.GetDatabasesOfTeacher(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, databasesOfTeacher)
}

func (h *HttpHandler) NewDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *HttpHandler) GetDatabase(c *gin.Context) {
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

func (h *HttpHandler) EditDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *HttpHandler) DeleteDatabase(c *gin.Context) {
	panic("implement me")
}

type DatabaseForm struct {
	ID            int32   `json:"id"`
	Name          string  `json:"name"`
	Data          string  `json:"data"`
	SchemaSVGPath *string `json:"schema_svg_path"`
	IsPublic      bool    `json:"is_public"`
}
