package databaseTemplates

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetDatabaseTemplates() (*[]model.DatabaseTemplateListing, error)
	NewDatabase(teacherID int32, name string) (*model.DatabaseTemplate, error)
	GetDatabase(databaseId int32, withData bool) (*model.DatabaseTemplate, error)
	EditDatabase(databaseId int32, sql string, imageData []byte) (*model.DatabaseTemplate, error)
	DeleteDatabase(databaseId int32)
}
