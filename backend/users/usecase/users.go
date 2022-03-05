package usecase

import (
	"errors"
	"fmt"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/users"
)

type usecase struct {
	repo users.Repository
}

func NewUsecase(repository users.Repository) users.Usecase {
	return &usecase{repository}
}

func (u usecase) SignUpTeacher(email, password string) error {
	teacher, err := u.repo.GetTeacherByEmail(email)
	if err != nil {
		return err
	}
	if teacher != nil {
		return fmt.Errorf("TEACHER WITH EMAIL \"%v\" already exists", email)
	}
	hashedPwd := u.repo.HashString(password)
	return u.repo.CreateTeacher(email, hashedPwd)
}

func (u usecase) LogInTeacher(email, password string) (*model.Teacher, error) {
	teacher, err := u.repo.GetTeacherByEmail(email)
	if err != nil {
		return nil, err
	}
	if teacher == nil {
		return nil, fmt.Errorf("EMAIL \"%v\" NOT FOUND", email)
	}
	correctPwd := u.repo.HashedStringEquals(password, teacher.Password)
	if !correctPwd {
		return nil, errors.New("PASSWORD NOT CORRECT")
	}
	return teacher, nil
}

func (u usecase) DeleteTeacher(requestID int32) error {
	teacher, err := u.repo.GetTeacherByID(requestID)
	if err != nil {
		return err
	}
	if teacher == nil {
		return errors.New("ALREADY DELETED THIS ACCOUNT")
	}
	if teacher.ID != requestID {
		return errors.New("CANNOT DELETE ANOTHER USER")
	}
	return u.repo.DeleteTeacher(requestID)
}

func (u usecase) LogInParticipant(accessCode string) (*model.Participant, error) {
	panic("Not implemented")
}
