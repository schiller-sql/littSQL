package usecase

import (
	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/model"
)

type eUsecase struct {
	repo databases.Repository
}

func NewUsecase(repo databases.Repository) databases.Usecase {
	return &eUsecase{repo}
}

func (e eUsecase) GetDatabasesOfTeacher(teacherID string) []*model.DatabaseListing {
	panic("implement me")
}

func (e eUsecase) GetDatabaseDetails(id string) *model.Database {
	panic("implement me")
}

func (e eUsecase) NewDatabase(teacherID string, name string) *model.Database {
	panic("implement me")
}

func (e eUsecase) EditDatabase(teacherId string, id string, sql string, imageData []byte) (*model.Database, error) {
	panic("implement me")
}

func (e eUsecase) DeleteDatabase(teacherId string, id string) error {
	panic("implement me")
}
