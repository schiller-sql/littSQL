package auth

import "github.com/schiller-sql/littSQL/model"

type Usecase interface {
	SignUpTeacher(email, password string) error
	LogInTeacher(email, password string) (*model.Teacher, error)
	DeleteTeacher(teacherID int32) error
	LogInParticipant(accessCode string) (*model.Participant, error)
	GetTeacherAccountDetails(teacherID int32) (*model.Teacher, error)
	GetParticipantAccountDetails(participantID int32) (*model.Participant, error)
}
