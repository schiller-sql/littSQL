package usecase

import (
	"fmt"

	"github.com/schiller-sql/littSQL/database_templates"
	"github.com/schiller-sql/littSQL/model"
)

type eUsecase struct {
	databaseTemplatesRepo databaseTemplates.Repository
}

func NewUsecase(databaseTemplatesRepo databaseTemplates.Repository) databaseTemplates.Usecase {
	return &eUsecase{databaseTemplatesRepo}
}

func (e eUsecase) GetDatabaseTemplates() (*[]model.DatabaseTemplateListing, error) {
	return e.databaseTemplatesRepo.GetDatabaseTemplates()
}

func (e eUsecase) NewDatabase(teacherID int32, name string) *model.DatabaseTemplate {
	panic("implement me")
}

func (e eUsecase) GetDatabaseDetails(databaseId int32) (*model.DatabaseTemplate, error) {
	database, err := e.databaseTemplatesRepo.GetDatabase(databaseId, true)
	if err != nil {
		return nil, err
	}
	if database == nil {
		return nil, fmt.Errorf("the database template with the id '%d' does not exist", databaseId)
	}
	return database, nil
}

func (e eUsecase) EditDatabase(teacherID int32, databaseId int32, sql string, imageData []byte) (*model.DatabaseTemplate, error) {
	panic("implement me")
}

func (e eUsecase) DeleteDatabase(teacherID int32, databaseId int32) error {
	panic("implement me")
}
