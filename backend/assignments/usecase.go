package assignments

import (
	"github.com/schiller-sql/littSQL/model"
	"time"
)

type Usecase interface {
	GetAssignmentsOfCourse(teacherID, courseID int32) (*[]model.AssignmentListing, error)
	GetAssignment(teacherID, courseID, assignmentID int32) (*model.Assignment, error)
	NewAssignment(teacherID, courseID int32, name string) (*model.Assignment, error)
	EditAssignment(
		teacherID,
		courseID,
		assignmentID int32,
		name string,
		comment *string,
		projectID *int32,
		finishedDate *time.Time,
		locked bool,
		answerConfig model.AssignmentAnswerConfig,
	) error
	EditAssignmentOrder(teacherID, courseID, assignmentID, newOrder int32) error
	DeleteAssignment(teacherID, courseID, assignmentID int32) error
}
