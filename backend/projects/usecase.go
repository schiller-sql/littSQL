package projects

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	// TODO: Maybe add something similar in the future:
	//CloneProject(teacherID int32, projectID int32, name string) (*model.Project, error)
	GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error)
	EditProject(
		projectID int32,
		teacherID int32,
		name string,
		documentationMd string,
		sql *string,
		tasks []model.Task,
	) error
	DeleteProject(teacherID int32, projectID int32) error
}
