package databases

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetDatabasesOfTeacher(teacherId int) (*[]model.DatabaseListing, error)
	NewDatabase(teacherID string, name string) (*model.Database, error)
	GetDatabase(id string, withDB bool) (*model.Database, error)
	EditDatabase(id string, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(id string)
}
