package projects

import (
	"database/sql"
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error)
	GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error)
	NewProject(teacherID int32, name string) (*model.Project, error)
	EditProject(
		ID int32,
		TeacherID int32,
		DatabaseID sql.NullInt64,
		Name string,
		DocumentationMd string,
		Tasks []model.Task,
	) error
	DeleteProject(teacherID int32, projectID int32) error

	//NewSampleDatabase(name string, schemaPictureFile []byte, sqliteDatabaseFile []byte) model.Database
	//NewProject(name string) model.Project
	//EditProject(projectID int32, databaseID sql.NullInt64, teacherID int32, documentationMd string, tasks []model.Task) model.Project
	//GetAllProjects(c *gin.Context)
	//GetProject(c *gin.Context)
	//EditProject(c *gin.Context)
	//DeleteProject(c *gin.Context)
	//NewProject(c *gin.Context)
}
