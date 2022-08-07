package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/assignments"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"github.com/schiller-sql/littSQL/helpers"
	"github.com/schiller-sql/littSQL/model"
	"net/http"
	"time"
)

type assignmentsHandler struct {
	usecase assignments.Usecase
}

func ConfigureHandler(courseGroup *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase assignments.Usecase) {
	handler := assignmentsHandler{usecase}

	group := courseGroup.Group("/assignments", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	group.GET("", handler.getAssignmentsOfCourse)
	group.POST("", handler.newAssignment)
	group.GET("/:assignment-id", handler.getAssignment)
	group.PUT("/:assignment-id", handler.editAssignment)
	group.POST("/:assignment-id/order", handler.editAssignmentOrder)
	group.DELETE("/:assignment-id", handler.deleteAssignment)
}

func (h *assignmentsHandler) getIDs(c *gin.Context) (teacherID, courseID int32, err error) {
	teacherID = helpers.GetJwtID(c)
	courseID, err = helpers.GetParamIDInt32(c)
	if err != nil {
		return
	}
	return
}

func (h *assignmentsHandler) getAssignment(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assignmentID, err := helpers.GetParamInt32(c, "assignment-id")
	if err != nil {
		return
	}
	assignmentsOfCourse, err := h.usecase.GetAssignment(teacherID, courseID, assignmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assignmentsOfCourse)
}

func (h *assignmentsHandler) getAssignmentsOfCourse(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assignmentsOfCourse, err := h.usecase.GetAssignmentsOfCourse(teacherID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assignmentsOfCourse)
}

func (h *assignmentsHandler) newAssignment(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newAssignmentForm NewAssignmentForm
	err = c.BindJSON(&newAssignmentForm)
	if err != nil {
		return
	}
	newAssignment, err := h.usecase.NewAssignment(
		teacherID,
		courseID,
		newAssignmentForm.Name,
		newAssignmentForm.Comment,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newAssignment)
}

func (h *assignmentsHandler) editAssignment(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assignmentID, err := helpers.GetParamInt32(c, "assignment-id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var form AssignmentEditForm
	err = c.BindJSON(&form)
	if err != nil {
		return
	}
	err = h.usecase.EditAssignment(
		teacherID,
		courseID,
		assignmentID,
		form.Name,
		form.Comment,
		form.ProjectID,
		form.FinishedDate,
		*form.Locked,
		form.AnswerConfig,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *assignmentsHandler) editAssignmentOrder(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assignmentID, err := helpers.GetParamInt32(c, "assignment-id")
	if err != nil {
		return
	}
	var newOrder int32
	err = c.BindJSON(&newOrder)
	if err != nil {
		return
	}
	err = h.usecase.EditAssignmentOrder(teacherID, courseID, assignmentID, newOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *assignmentsHandler) deleteAssignment(c *gin.Context) {
	teacherID, courseID, err := h.getIDs(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assignmentID, err := helpers.GetParamInt32(c, "assignment-id")
	if err != nil {
		return
	}
	err = h.usecase.DeleteAssignment(teacherID, courseID, assignmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type NewAssignmentForm struct {
	Name    string  `json:"name" binding:"required"`
	Comment *string `json:"comment"`
}

type AssignmentEditForm struct {
	Name         string                       `json:"name" binding:"required"`
	Comment      *string                      `json:"comment"`
	ProjectID    *int32                       `json:"project_id"`
	FinishedDate *time.Time                   `json:"finished_date"`
	AnswerConfig model.AssignmentAnswerConfig `json:"answer_config" binding:"required"`
	Locked       *bool                        `json:"locked" binding:"required"`
}
