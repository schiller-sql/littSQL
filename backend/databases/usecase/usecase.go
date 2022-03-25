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

func (e eUsecase) GetDatabasesOfTeacher(teacherID int32) []*model.DatabaseListing {
	panic("implement me")
}

func (e eUsecase) NewDatabase(teacherID int32, name string) *model.Database {
	panic("implement me")
}

func (e eUsecase) GetDatabaseDetails(id int32) (*model.Database, error) {
	return e.repo.GetDatabase(id, true)
}

func (e eUsecase) EditDatabase(teacherId int32, id int32, sql string, imageData []byte) (*model.Database, error) {
	panic("implement me")
}

func (e eUsecase) DeleteDatabase(teacherId int32, id int32) error {
	panic("implement me")
}
