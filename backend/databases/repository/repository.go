package repository

import (
	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/model"
	"gorm.io/gorm"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) databases.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetDatabasesOfTeacher(teacherId int) (*[]model.DatabaseListing, error) {
	panic("implement me")
}

func (e eRepository) NewDatabase(teacherID string, name string) (*model.Database, error) {
	panic("implement me")
}

func (e eRepository) GetDatabase(id string, withDB bool) (*model.Database, error) {
	panic("implement me")
}

func (e eRepository) EditDatabase(id string, sql string, imageData []byte) (*model.Database, error) {
	panic("implement me")
}

func (e eRepository) DeleteDatabase(id string) {
	panic("implement me")
}
