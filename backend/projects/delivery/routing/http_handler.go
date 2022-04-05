package routing

import (
	"github.com/gin-gonic/gin"
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"github.com/schiller-sql/littSQL/helpers"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
	"net/http"
)

type projectsHandler struct {
	usecase projects.Usecase
}

func ConfigureHandler(r *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase projects.Usecase) {
	handler := projectsHandler{usecase}

	group := r.Group("/projects", authMiddleware.JwtHandler, authMiddleware.IsTeacherValidator)

	group.GET("/", handler.getProjectsOfTeacher)
	group.POST("/", handler.newProject)
	group.GET("/:id", handler.getProject)
	group.PUT("/:id", handler.editProject)
	group.DELETE("/:id", handler.deleteProject)
}

func (h *projectsHandler) getProjectsOfTeacher(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	projectsOfTeacher, err := h.usecase.GetProjectsOfTeacher(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectsOfTeacher)
}

func (h *projectsHandler) editProject(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	projectID, err := helpers.GetParamIDInt32(c)
	if err != nil {
		return
	}
	var projectEditForm ProjectEditForm
	err = c.BindJSON(&projectEditForm)
	if err != nil {
		return
	}
	tasks := taskFormsToTasks(projectID, projectEditForm.Tasks)
	err = h.usecase.EditProject(
		projectID,
		teacherID,
		projectEditForm.Name,
		projectEditForm.DocumentationMd,
		projectEditForm.DbSQL,
		tasks,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *projectsHandler) getProject(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	projectID, err := helpers.GetParamIDInt32(c)
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

func (h *projectsHandler) deleteProject(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
	projectID, err := helpers.GetParamIDInt32(c)
	if err != nil {
		return
	}
	err = h.usecase.DeleteProject(teacherID, projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *projectsHandler) newProject(c *gin.Context) {
	teacherID := helpers.GetJwtID(c)
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

func projectToProjectReturnForm(project model.Project) ProjectReturnForm {
	taskForms := make([]TaskForm, len(project.Tasks))
	for i := 0; i < len(taskForms); i++ {
		for j := 0; j < len(taskForms); j++ {
			if project.Tasks[j].Number == int32(i) {
				task := project.Tasks[j]
				questions := task.Questions
				questionForms := make([]QuestionForm, len(questions))
				for k := 0; k < len(questionForms); k++ {
					for l := 0; l < len(questionForms); l++ {
						if task.Questions[l].Number == int32(k) {
							question := questions[l]
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
	return ProjectReturnForm{
		ID:       project.ID,
		IsPublic: project.OwnerID == nil,
		ProjectEditForm: ProjectEditForm{
			DbSQL:           project.DbSQL,
			Name:            project.Name,
			DocumentationMd: project.DocumentationMd,
			Tasks:           taskForms,
		},
	}
}

func taskFormsToTasks(projectID int32, taskForms []TaskForm) []model.Task {
	tasks := make([]model.Task, len(taskForms))
	for i := 0; i < len(taskForms); i++ {
		taskNumber := int32(i)
		questionForms := taskForms[i].Questions
		questions := make([]model.Question, len(questionForms))
		for j := 0; j < len(questions); j++ {
			questionNumber := int32(j)
			questions[j] = model.Question{
				ProjectID:  projectID,
				TaskNumber: taskNumber,
				Number:     questionNumber,
				Question:   questionForms[j].Question,
				Type:       questionForms[j].Type,
				Solution:   questionForms[j].Solution,
			}
		}
		tasks[i] = model.Task{
			ProjectID:   projectID,
			Number:      taskNumber,
			Description: taskForms[i].Description,
			IsVoluntary: taskForms[i].IsVoluntary,
			Questions:   questions,
		}
	}
	return tasks
}

type ProjectReturnForm struct {
	ID       int32 `json:"id"`
	IsPublic bool  `json:"is_public"`
	ProjectEditForm
}

type ProjectEditForm struct {
	DbSQL           *string    `json:"sql"`
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
