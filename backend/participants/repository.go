package participants

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetParticipantsOfCourse(courseID int32) (*[]model.ParticipantListing, error)
	NewParticipant(courseID int32, name *string) (*model.Participant, error)
	GetParticipant(participantID int32) (*model.Participant, error)
	EditParticipant(participantID int32, name *string) error
	RefreshParticipantAccessCode(participantID int32) (string, error)
	DeleteParticipant(participantID int32) error
}
