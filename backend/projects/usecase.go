package projects

import (
	"database/sql"
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	// TODO: Sort by: is not public; name
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error)
	EditProject(
		ID int32,
		TeacherID int32,
		DatabaseID sql.NullInt64,
		Name string,
		DocumentationMd string,
		Tasks []model.Task,
	) error
	DeleteProject(teacherID int32, projectID int32) error
}
