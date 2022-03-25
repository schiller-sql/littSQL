package routing

import (
	"database/sql"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
	"net/http"
	"strconv"
)

type projectsHandler struct {
	usecase projects.Usecase
}

func ConfigureHandler(r *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware, usecase projects.Usecase) {
	handler := projectsHandler{usecase}

	group := r.Group("/projects", jwtMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if !claims["is_teacher"].(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You have to be a teacher to access this resource"})
		}
	})

	group.GET("/", handler.getProjectsOfTeacher)
	group.POST("/", handler.newProject)
	group.GET("/:id", handler.getProject)
	group.PUT("/:id", handler.editProject)
	group.DELETE("/:id", handler.deleteProject)
}

func getTeacherIDHelper(c *gin.Context) int32 {
	id, _ := c.Get("id")
	return id.(int32)
}

func getProjectIDHelper(c *gin.Context) (int32, error) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}
	return int32(projectID), nil
}

func (h *projectsHandler) getProjectsOfTeacher(c *gin.Context) {
	teacherID := getTeacherIDHelper(c)
	projectsOfTeacher, err := h.usecase.GetProjectsOfTeacher(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectsOfTeacher)
}

func (h *projectsHandler) getProject(c *gin.Context) {
	teacherID := getTeacherIDHelper(c)
	projectID, err := getProjectIDHelper(c)
	if err != nil {
		return
	}
	project, err := h.usecase.GetProjectDetails(teacherID, projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projectForm := projectToProjectReturnForm(*project)
	c.JSON(http.StatusOK, projectForm)
}

func (h *projectsHandler) editProject(c *gin.Context) {
	// TODO: check that database_id exists
	teacherID := getTeacherIDHelper(c)
	projectID, err := getProjectIDHelper(c)
	if err != nil {
		return
	}
	var projectEditForm ProjectEditForm
	err = c.BindJSON(&projectEditForm)
	if err != nil {
		return
	}
	databaseID := sql.NullInt64{Valid: true}
	if projectEditForm.DatabaseID != nil {
		databaseID.Int64 = int64(*projectEditForm.DatabaseID)
	} else {
		databaseID.Valid = false
	}
	tasks := taskFormsToTasks(projectID, projectEditForm.Tasks)
	err = h.usecase.EditProject(
		projectID,
		teacherID,
		databaseID,
		projectEditForm.Name,
		projectEditForm.DocumentationMd,
		tasks,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *projectsHandler) deleteProject(c *gin.Context) {
	teacherID := getTeacherIDHelper(c)
	projectID, err := getProjectIDHelper(c)
	if err != nil {
		return
	}
	err = h.usecase.DeleteProject(teacherID, projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *projectsHandler) newProject(c *gin.Context) {
	teacherID := getTeacherIDHelper(c)
	var newProjectForm NewProjectForm
	err := c.BindJSON(&newProjectForm)
	if err != nil {
		return
	}
	newProject, err := h.usecase.NewProject(teacherID, newProjectForm.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectToProjectReturnForm(*newProject))
}

type NewProjectForm struct {
	Name string `json:"name" binding:"required"`
}

// TODO: Possibly conversions not working correctly
func projectToProjectReturnForm(project model.Project) ProjectReturnForm {
	taskForms := make([]TaskForm, len(project.Tasks))
	for i := 0; i < len(taskForms); i++ {
		for j := 0; j < len(taskForms); j++ {
			if project.Tasks[j].Number == int32(i) {
				task := project.Tasks[j]
				questionForms := make([]QuestionForm, len(task.Questions))
				for k := 0; k < len(questionForms); k++ {
					for l := 0; l < len(questionForms); l++ {
						if task.Questions[l].TaskNumber == int32(k) {
							question := task.Questions[l]
							questionForms[k] = QuestionForm{
								Question: question.Question,
								Type:     question.Type,
								Solution: question.Solution,
							}
							break
						}
					}
				}
				taskForms[i] = TaskForm{Description: task.Description, IsVoluntary: task.IsVoluntary, Questions: questionForms}
				break
			}
		}
	}
	var databaseID *int32
	if project.DatabaseID.Valid {
		id := int32(project.DatabaseID.Int64)
		databaseID = &id
	}
	return ProjectReturnForm{
		ID: project.ID,
		ProjectEditForm: ProjectEditForm{
			DatabaseID:      databaseID,
			Name:            project.Name,
			DocumentationMd: project.DocumentationMd,
			Tasks:           taskForms,
		},
	}
}

func taskFormsToTasks(projectId int32, taskForms []TaskForm) []model.Task {
	tasks := make([]model.Task, len(taskForms))
	for i := 0; i < len(taskForms); i++ {
		taskNumber := int32(i)
		questionForms := taskForms[i].Questions
		questions := make([]model.Question, len(questionForms))
		for j := 0; j < len(questions); j++ {
			questionNumber := int32(j)
			questions[j] = model.Question{
				ProjectID:  projectId,
				TaskNumber: taskNumber,
				Number:     questionNumber,
				Question:   questionForms[j].Question,
				Type:       questionForms[j].Type,
				Solution:   questionForms[j].Solution,
			}
		}
		tasks[i] = model.Task{
			ProjectID:   projectId,
			Number:      taskNumber,
			Description: taskForms[i].Description,
			IsVoluntary: taskForms[i].IsVoluntary,
			Questions:   questions,
		}
	}
	return tasks
}

type ProjectReturnForm struct {
	ID int32 `json:"id"`
	ProjectEditForm
}

type ProjectEditForm struct {
	DatabaseID      *int32     `json:"database_id"`
	Name            string     `json:"name" binding:"required"`
	DocumentationMd string     `json:"documentation_md"`
	Tasks           []TaskForm `json:"tasks" binding:"required"`
}

type TaskForm struct {
	Description string         `json:"description" binding:"required"`
	IsVoluntary bool           `json:"is_voluntary" binding:"required"`
	Questions   []QuestionForm `json:"questions" binding:"required"`
}

type QuestionForm struct {
	Question string `json:"question" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Solution string `json:"solution" binding:"required"`
}
