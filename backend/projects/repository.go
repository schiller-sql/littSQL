package projects

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	// TODO: Sort by: is not public; name
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	GetProjectSuperficial(projectID int32) (*model.ProjectSuperficial, error)
	GetProject(projectID int32) (*model.Project, error)
	SaveEditedProject(editedProject *model.Project) error
	DeleteProject(projectID int32) error
}
