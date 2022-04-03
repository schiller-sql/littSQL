package routing

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"github.com/schiller-sql/littSQL/courses"
	"github.com/schiller-sql/littSQL/helpers"
	"net/http"
)

type coursesHandler struct {
	usecase courses.Usecase
}

func ConfigureHandler(r *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase courses.Usecase) {
	handler := coursesHandler{usecase}

	group := r.Group("/courses", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	group.GET("/", handler.getCoursesOfTeacher)
	group.POST("/", handler.newCourse)
	group.GET("/:id", handler.getCourse)
	group.PUT("/:id", handler.editCourse)
	group.DELETE("/:id", handler.deleteCourse)
}

func (h *coursesHandler) getCoursesOfTeacher(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	coursesOfTeacher, err := h.usecase.GetCoursesOfTeacher(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coursesOfTeacher)
}

func (h *coursesHandler) newCourse(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	var newCourseForm CourseEditForm
	err := c.BindJSON(&newCourseForm)
	if err != nil {
		return
	}
	newCourse, err := h.usecase.NewCourse(teacherID, newCourseForm.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, CourseReturnForm{ID: newCourse.ID, Name: newCourse.Name})
}

func (h *coursesHandler) getCourse(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	courseID, err := helpers.GetParamID(c)
	if err != nil {
		return
	}
	course, err := h.usecase.GetCourseDetails(teacherID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, CourseReturnForm{ID: course.ID, Name: course.Name})
}

func (h *coursesHandler) editCourse(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	courseID, err := helpers.GetParamID(c)
	if err != nil {
		return
	}
	var form CourseEditForm
	err = c.BindJSON(&form)
	if err != nil {
		return
	}
	err = h.usecase.EditCourse(courseID, teacherID, form.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *coursesHandler) deleteCourse(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	courseID, err := helpers.GetParamID(c)
	if err != nil {
		return
	}
	err = h.usecase.DeleteCourse(teacherID, courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type CourseReturnForm struct {
	ID   int32  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CourseEditForm struct {
	Name string `json:"name" binding:"required"`
}
