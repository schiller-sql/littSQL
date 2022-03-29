package repository

import (
	"github.com/schiller-sql/littSQL/database_templates"
	"github.com/schiller-sql/littSQL/model"
	"gorm.io/gorm"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) databaseTemplates.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetDatabaseTemplates() (*[]model.DatabaseTemplateListing, error) {
	var databaseTemplatesOfTeacher []model.DatabaseTemplateListing
	result := e.DB.Table("database_templates").Find(&databaseTemplatesOfTeacher)
	return &databaseTemplatesOfTeacher, result.Error
}

func (e eRepository) NewDatabase(teacherID int32, name string) (*model.DatabaseTemplate, error) {
	panic("implement me")
}

func (e eRepository) GetDatabase(databaseId int32, withData bool) (*model.DatabaseTemplate, error) {
	var database model.DatabaseTemplate
	query := e.DB
	if !withData {
		query = query.Omit("sql")
	}
	result := query.Find(&database, databaseId)
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &database, nil
}

func (e eRepository) EditDatabase(databaseId int32, sql string, imageData []byte) (*model.DatabaseTemplate, error) {
	panic("implement me")
}

func (e eRepository) DeleteDatabase(databaseId int32) {
	panic("implement me")
}
