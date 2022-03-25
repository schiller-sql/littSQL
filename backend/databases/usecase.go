package databases

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetDatabasesOfTeacher(teacherID int32) (*[]model.DatabaseListing, error)
	GetDatabaseDetails(id int32) (*model.Database, error)
	NewDatabase(teacherID int32, name string) *model.Database
	EditDatabase(teacherID int32, id int32, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(teacherID int32, id int32) error
}
