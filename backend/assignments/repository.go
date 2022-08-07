package assignments

import (
	"github.com/schiller-sql/littSQL/model"
	"time"
)

type Repository interface {
	GetAssignmentsOfCourse(courseID int32) (*[]model.Assignment, error)
	GetAmountOfAssignmentsOfCourse(courseID int32) (int32, error)
	NewAssignment(courseID int32, name string) (*model.Assignment, error)
	GetAssignment(assignmentID int32) (*model.Assignment, error)
	EditAssignment(
		assignmentID int32,
		name string,
		comment *string,
		projectID *int32,
		finishedDate *time.Time,
		locked bool,
		answerConfig model.AssignmentAnswerConfig,
	) error
	EditAssignmentOrder(courseID int32, assignmentID int32, newOrder int32) error
	DeleteAssignment(assignmentID int32) error
}
