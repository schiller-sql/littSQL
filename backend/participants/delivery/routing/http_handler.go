package routing

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"github.com/schiller-sql/littSQL/helpers"
	"github.com/schiller-sql/littSQL/participants"
	"net/http"
)

type participantsHandler struct {
	usecase participants.Usecase
}

func ConfigureHandler(courseGroup *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase participants.Usecase) {
	handler := participantsHandler{usecase}

	group := courseGroup.Group("/participants", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	group.GET("", handler.getParticipantsOfCourse)
	group.POST("", handler.newParticipant)
	group.PUT("/:participant-id", handler.editParticipant)
	group.PUT("/:participant-id/refresh-access-code", handler.refreshParticipantAccessCode)
	group.DELETE("/:participant-id", handler.deleteParticipant)
}

func (h *participantsHandler) getIDs(c *gin.Context) (teacherID, courseID int32, err error) {
	teacherID = helpers.GetJwtID(c)
	courseID, err = helpers.GetParamIDInt32(c)
	if err != nil {
		return
	}
	return
}

func (h *participantsHandler) getParticipantsOfCourse(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	participantsOfCourse, err := h.usecase.GetParticipantsOfCourse(teacherID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, participantsOfCourse)
}

func (h *participantsHandler) newParticipant(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newParticipantForm ParticipantEditForm
	err = c.BindJSON(&newParticipantForm)
	if err != nil {
		return
	}
	newParticipant, err := h.usecase.NewParticipant(teacherID, courseID, newParticipantForm.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ParticipantReturnForm{
		ID:         newParticipant.ID,
		Name:       newParticipant.Name,
		AccessCode: newParticipant.AccessCode,
	})
}

func (h *participantsHandler) editParticipant(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	participantID, err := helpers.GetParamInt32(c, "participant-id")
	if err != nil {
		return
	}
	var form ParticipantEditForm
	err = c.BindJSON(&form)
	if err != nil {
		return
	}
	err = h.usecase.EditParticipant(teacherID, courseID, participantID, form.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *participantsHandler) refreshParticipantAccessCode(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	participantID, err := helpers.GetParamInt32(c, "participant-id")
	if err != nil {
		return
	}
	newAccessCode, err := h.usecase.RefreshParticipantAccessCode(teacherID, courseID, participantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_code": newAccessCode})
}

func (h *participantsHandler) deleteParticipant(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	participantID, err := helpers.GetParamInt32(c, "participant-id")
	if err != nil {
		return
	}
	err = h.usecase.DeleteParticipant(teacherID, courseID, participantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type ParticipantReturnForm struct {
	ID         int32   `json:"id"`
	Name       *string `json:"name"`
	AccessCode string  `json:"access_code"`
}

type ParticipantEditForm struct {
	Name *string `json:"name"`
}
