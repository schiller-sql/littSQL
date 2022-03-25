package databases

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetDatabasesOfTeacher(teacherID string) []*model.DatabaseListing
	GetDatabaseDetails(id string) *model.Database
	NewDatabase(teacherID string, name string) *model.Database
	EditDatabase(teacherId string, id string, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(teacherId string, id string) error
}
