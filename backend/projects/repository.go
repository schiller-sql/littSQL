package projects

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	// TODO: Sort by: is not public; name
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	GetProject(projectID int32, tasks bool) (*model.Project, error)
	SaveEditedProject(editedProject *model.Project) error
	DeleteProject(projectID int32) error
}
