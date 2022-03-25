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

func (e eRepository) GetDatabasesOfTeacher(teacherID int32) (*[]model.DatabaseListing, error) {
	var databasesOfTeacher []model.DatabaseListing
	result := e.DB.Raw(
		"select id, name, owner_id is null as is_public from databases where owner_id is null or owner_id = ?",
		teacherID,
	).Find(&databasesOfTeacher)
	return &databasesOfTeacher, result.Error
}

func (e eRepository) NewDatabase(teacherID int32, name string) (*model.Database, error) {
	panic("implement me")
}

func (e eRepository) GetDatabase(id int32, withData bool) (*model.Database, error) {
	var database model.Database
	query := e.DB
	if !withData {
		query = query.Omit("data")
	}
	result := query.Find(&database, id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &database, nil
}

func (e eRepository) EditDatabase(id int32, sql string, imageData []byte) (*model.Database, error) {
	panic("implement me")
}

func (e eRepository) DeleteDatabase(id int32) {
	panic("implement me")
}
