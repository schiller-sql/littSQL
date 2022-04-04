package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	databaseTemplates "github.com/schiller-sql/littSQL/database_templates"
	"github.com/schiller-sql/littSQL/helpers"
)

type databaseTemplatesHandler struct {
	usecase databaseTemplates.Usecase
}

func ConfigureHandler(r *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase databaseTemplates.Usecase) {
	handler := databaseTemplatesHandler{usecase}

	group := r.Group("/database-templates", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	// TODO: implement other methods, or let it be
	group.GET("/", handler.getDatabaseTemplates)
	//group.POST("/", handler.newDatabase)
	group.GET("/:id", handler.getDatabase)
	//group.PUT("/:id", handler.editDatabase)
	//group.DELETE("/:id", handler.deleteDatabase)
}

func (h *databaseTemplatesHandler) getDatabaseTemplates(c *gin.Context) {
	databaseTemplatesOfTeacher, err := h.usecase.GetDatabaseTemplates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, databaseTemplatesOfTeacher)
}

func (h *databaseTemplatesHandler) newDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databaseTemplatesHandler) getDatabase(c *gin.Context) {
	databaseID, err := helpers.GetParamID(c)
	if err != nil {
		return
	}
	database, err := h.usecase.GetDatabaseDetails(databaseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, database)
}

func (h *databaseTemplatesHandler) editDatabase(c *gin.Context) {
	panic("implement me")
}

func (h *databaseTemplatesHandler) deleteDatabase(c *gin.Context) {
	panic("implement me")
}
