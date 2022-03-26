package databases

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetDatabasesOfTeacher(teacherID int32) (*[]model.DatabaseListing, error)
	NewDatabase(teacherID int32, name string) (*model.Database, error)
	GetDatabase(databaseId int32, withData bool) (*model.Database, error)
	EditDatabase(databaseId int32, sql string, imageData []byte) (*model.Database, error)
	DeleteDatabase(databaseId int32)
}
