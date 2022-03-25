package databases

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetDatabasesOfTeacher(teacherID int32) []*model.DatabaseListing
	GetDatabaseDetails(id int32) (*model.Database, error)
	NewDatabase(teacherID int32, name string) *model.Database
	EditDatabase(teacherId int32, id int32, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(teacherId int32, id int32) error
}
