package users

import "github.com/schiller-sql/littSQL/model"

type Usecase interface {
	SignUpTeacher(email, password string) error
	LogInTeacher(email, password string) (*model.Teacher, error)
	DeleteTeacher(teacherID int32) error
	LogInParticipant(accessCode string) (*model.Participant, error)
}
