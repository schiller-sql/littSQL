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
		return fmt.Errorf("TEACHER WITH EMAIL \"%v\" ALREADY EXISTS", email)
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
	participant, err := u.repo.GetParticipantByAccessCode(accessCode)
	if err != nil {
		return nil, err
	}
	if participant == nil {
		return nil, errors.New("ACCESS CODE DOES NOT EXIST")
	}
	return participant, nil
}

func (u usecase) GetTeacherAccountDetails(teacherID int32) (*model.Teacher, error) {
	teacher, err := u.repo.GetTeacherByID(teacherID)
	if err != nil {
		return nil, err
	}
	if teacher == nil {
		return nil, errors.New("ACCOUNT NOT FOUND")
	}
	return teacher, nil
}

func (u usecase) GetParticipantAccountDetails(participantID int32) (*model.Participant, error) {
	participant, err := u.repo.GetParticipantByID(participantID)
	if err != nil {
		return nil, err
	}
	if participant == nil {
		return nil, errors.New("ACCOUNT NOT FOUND")
	}
	return participant, nil
}
