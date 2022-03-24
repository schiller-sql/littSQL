package databases

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetDatabasesOfTeacher(teacherId int) ([]model.Database, error)
	GetDatabase(id string) (model.Database, error)
	NewDatabase(name string) (model.Database, error)
	EditDatabase(id string, editedDatabase model.Database) (model.Database, error)
	DeleteDatabase(id string)
}
