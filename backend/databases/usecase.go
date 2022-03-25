package databases

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	// TODO: Sort by: is not public; name
	GetDatabasesOfTeacher(teacherID int32) (*[]model.DatabaseListing, error)
	NewDatabase(teacherID int32, name string) *model.Database
	GetDatabaseDetails(id int32) (*model.Database, error)
	EditDatabase(teacherID int32, id int32, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(teacherID int32, id int32) error
}
