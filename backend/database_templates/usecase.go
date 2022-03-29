package databaseTemplates

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	// TODO: Sort by: is not public; name
	GetDatabaseTemplates() (*[]model.DatabaseTemplateListing, error)
	NewDatabase(teacherID int32, name string) *model.DatabaseTemplate
	GetDatabaseDetails(databaseId int32) (*model.DatabaseTemplate, error)
	EditDatabase(teacherID int32, databaseId int32, sql string, imageData []byte) (*model.DatabaseTemplate, error)
	DeleteDatabase(teacherID int32, databaseId int32) error
}
