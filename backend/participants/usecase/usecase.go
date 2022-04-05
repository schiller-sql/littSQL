package usecase

import (
	"fmt"
	"github.com/schiller-sql/littSQL/courses"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/participants"
)

type eUsecase struct {
	participantsRepo participants.Repository
	coursesRepo      courses.Repository
}

func NewUsecase(participantsRepo participants.Repository, coursesRepo courses.Repository) participants.Usecase {
	return &eUsecase{participantsRepo, coursesRepo}
}

func (u eUsecase) checkCourseBelongsToTeacher(courseID, teacherID int32) error {
	course, err := u.coursesRepo.GetCourse(courseID)
	if course == nil {
		return fmt.Errorf("course with id '%d' does not exist", courseID)
	}
	if err != nil {
		return err
	}
	if course.TeacherID != teacherID {
		return fmt.Errorf("course with id '%d', does not belong to you", courseID)
	}
	return nil
}

func (u eUsecase) checkParticipantBelongsToCourse(participantID, courseID int32) error {
	participant, err := u.participantsRepo.GetParticipant(participantID)
	if participant == nil {
		return fmt.Errorf("student with id '%d' does not exist", participantID)
	}
	if err != nil {
		return err
	}
	if participant.CourseID != courseID {
		return fmt.Errorf(
			"participant '%d' with id '%v', does not belong to course with id '%d'",
			participant.Name,
			participant,
			courseID,
		)
	}
	return nil
}

func (u eUsecase) GetParticipantsOfCourse(teacherID, courseID int32) (*[]model.ParticipantListing, error) {
	if err := u.checkCourseBelongsToTeacher(courseID, teacherID); err != nil {
		return nil, err
	}
	return u.participantsRepo.GetParticipantsOfCourse(courseID)
}

func (u eUsecase) NewParticipant(teacherID, courseID int32, name *string) (*model.Participant, error) {
	if err := u.checkCourseBelongsToTeacher(courseID, teacherID); err != nil {
		return nil, err
	}
	return u.participantsRepo.NewParticipant(courseID, name)
}

func (u eUsecase) EditParticipant(teacherID, courseID, participantID int32, name *string) error {
	if err := u.checkCourseBelongsToTeacher(courseID, teacherID); err != nil {
		return err
	}
	if err := u.checkParticipantBelongsToCourse(participantID, courseID); err != nil {
		return err
	}
	return u.participantsRepo.EditParticipant(participantID, name)
}

func (u eUsecase) RefreshParticipantAccessCode(teacherID, courseID, participantID int32) (string, error) {
	if err := u.checkCourseBelongsToTeacher(courseID, teacherID); err != nil {
		return "", err
	}
	if err := u.checkParticipantBelongsToCourse(participantID, courseID); err != nil {
		return "", err
	}
	return u.participantsRepo.RefreshParticipantAccessCode(participantID)
}

func (u eUsecase) DeleteParticipant(teacherID, courseID, participantID int32) error {
	if err := u.checkCourseBelongsToTeacher(courseID, teacherID); err != nil {
		return err
	}
	if err := u.checkParticipantBelongsToCourse(participantID, courseID); err != nil {
		return err
	}
	return u.participantsRepo.DeleteParticipant(participantID)
}
