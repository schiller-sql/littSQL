package participants

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetParticipantsOfCourse(teacherID int32, courseID int32) (*[]model.ParticipantListing, error)
	NewParticipant(teacherID int32, courseID int32, name *string) (*model.Participant, error)
	EditParticipant(teacherID int32, courseID int32, participantID int32, name *string) error
	RefreshParticipantAccessCode(teacherID int32, courseID int32, participantID int32) (string, error)
	DeleteParticipant(teacherID int32, courseID int32, participantID int32) error
}
