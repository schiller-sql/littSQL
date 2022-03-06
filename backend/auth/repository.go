package auth

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	DeleteTeacher(id int32) error
	CreateTeacher(email, password string) error
	GetTeacherByEmail(email string) (*model.Teacher, error)
	GetTeacherByID(teacherID int32) (*model.Teacher, error)
	GetParticipantByAccessCode(accessCode string) (*model.Participant, error)
	GetParticipantByID(participantID int32) (*model.Participant, error)
	HashString(s string) string
	HashedStringEquals(s, hashedS string) bool
}
