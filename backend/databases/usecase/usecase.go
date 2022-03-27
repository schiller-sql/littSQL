package usecase

import (
	"fmt"

	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/model"
)

type eUsecase struct {
	databasesRepo databases.Repository
}

func NewUsecase(databasesRepo databases.Repository) databases.Usecase {
	return &eUsecase{databasesRepo}
}

func (e eUsecase) GetDatabasesOfTeacher(teacherID int32) (*[]model.DatabaseListing, error) {
	return e.databasesRepo.GetDatabasesOfTeacher(teacherID)
}

func (e eUsecase) NewDatabase(teacherID int32, name string) *model.Database {
	panic("implement me")
}

func (e eUsecase) GetDatabaseDetails(databaseId int32) (*model.Database, error) {
	database, err := e.databasesRepo.GetDatabase(databaseId, true)
	if err != nil {
		return nil, err
	}
	if database == nil {
		return nil, fmt.Errorf("The database with the id '%d' does not exist", databaseId)
	}
	return database, nil
}

func (e eUsecase) EditDatabase(teacherID int32, databaseId int32, sql string, imageData []byte) (*model.Database, error) {
	panic("implement me")
}

func (e eUsecase) DeleteDatabase(teacherID int32, databaseId int32) error {
	panic("implement me")
}
