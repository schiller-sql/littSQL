package repository

import (
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/participants"
	"gorm.io/gorm"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) participants.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetParticipantsOfCourse(courseID int32) (*[]model.ParticipantListing, error) {
	var participantsOfCourse []model.ParticipantListing
	err := e.DB.Model(&model.Participant{}).Order("Upper(name)").Find(&participantsOfCourse, &model.Participant{CourseID: courseID}).Error
	if err != nil {
		return nil, err
	}
	return &participantsOfCourse, nil
}

func (e eRepository) NewParticipant(courseID int32, name *string) (*model.Participant, error) {
	participant := model.Participant{Name: name, CourseID: courseID}
	err := e.DB.Create(&participant).Error
	return &participant, err
}

func (e eRepository) GetParticipant(participantID int32) (*model.Participant, error) {
	var participant model.Participant
	result := e.DB.Find(&participant, participantID)
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &participant, nil
}

func (e eRepository) EditParticipant(participantID int32, name *string) error {
	return e.DB.Omit("course_id, access_code").Save(&model.Participant{ID: participantID, Name: name}).Error
}

func (e eRepository) RefreshParticipantAccessCode(participantID int32) (string, error) {
	var accessCode string
	err := e.DB.Raw(
		"update participants set access_code=utils.random_string(6) where id = ? returning access_code",
		participantID,
	).Find(&accessCode).Error
	if err != nil {
		return "", err
	}
	return accessCode, nil
}

func (e eRepository) DeleteParticipant(participantID int32) error {
	return e.DB.Delete(&model.Participant{}, participantID).Error
}
