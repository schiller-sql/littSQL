package repository

import (
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type eRepository struct {
	Cost int
	DB   *gorm.DB
}

func NewRepository(db *gorm.DB, cost int) users.Repository {
	return &eRepository{Cost: cost, DB: db}
}

func (e eRepository) DeleteTeacher(id int32) error {
	return e.DB.Delete(&model.Teacher{ID: id}).Error
}

func (e eRepository) CreateTeacher(email, hashedPassword string) error {
	newTeacher := model.Teacher{Email: email, Password: hashedPassword}
	return e.DB.Create(&newTeacher).Error
}

func (e eRepository) GetTeacherByEmail(email string) (*model.Teacher, error) {
	var teacher model.Teacher
	result := e.DB.Find(&teacher, &model.Teacher{Email: email})
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &teacher, nil
}

func (e eRepository) GetTeacherByID(teacherID int32) (*model.Teacher, error) {
	var teacher model.Teacher
	err := e.DB.Find(&teacher, teacherID).Error
	if teacher.Email == "" {
		return nil, nil
	}
	return &teacher, err
}

func (e eRepository) GetParticipantByAccessCode(accessCode string) (*model.Participant, error) {
	var participant model.Participant
	result := e.DB.Find(&participant, &model.Participant{AccessCode: accessCode})
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &participant, nil
}

func (e eRepository) GetParticipantByID(participantID int32) (*model.Participant, error) {
	var participant model.Participant
	err := e.DB.Find(&participant, participantID).Error
	if participant.AccessCode == "" {
		return nil, nil
	}
	return &participant, err
}

func (e eRepository) HashString(s string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(s), e.Cost)
	return string(password)
}

func (e eRepository) HashedStringEquals(s, hashedS string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedS), []byte(s))
	return err == nil
}
