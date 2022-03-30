package projects

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	GetProject(projectID int32, fillTasks bool) (*model.Project, error)
	SaveEditedProject(editedProject *model.Project) error
	DeleteProject(projectID int32) error
	SaveEditedProjectSqlData(projectID int32, data *[]byte) error
}
