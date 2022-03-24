package projects

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetProject(projectID int32, tasks bool) (*model.Project, error)
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	SaveEditedProject(editedProject *model.Project) error
	DeleteProject(projectID int32) error
}
