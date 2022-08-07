package usecase

import (
	"fmt"
	"github.com/schiller-sql/littSQL/assignments"
	"github.com/schiller-sql/littSQL/courses"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
	"time"
)

type eUsecase struct {
	assignmentsRepo assignments.Repository
	coursesRepo     courses.Repository
	projectsRepo    projects.Repository
}

func NewUsecase(
	participantsRepo assignments.Repository,
	coursesRepo courses.Repository,
	projectsRepo projects.Repository,
) assignments.Usecase {
	return &eUsecase{participantsRepo, coursesRepo, projectsRepo}
}

func (e eUsecase) checkCourseFromTeacher(teacherID, courseID int32) error {
	course, err := e.coursesRepo.GetCourse(courseID)
	if err != nil {
		return err
	}
	if course == nil {
		return fmt.Errorf("course with id %v does not exist", courseID)
	}
	if course.TeacherID != teacherID {
		return fmt.Errorf("course with id %v does not belong to your account", courseID)
	}
	return nil
}

func (e eUsecase) GetAssignmentsOfCourse(teacherID, courseID int32) (*[]model.Assignment, error) {
	err := e.checkCourseFromTeacher(teacherID, courseID)
	if err != nil {
		return nil, err
	}
	return e.assignmentsRepo.GetAssignmentsOfCourse(courseID)
}

func (e eUsecase) GetAssignment(teacherID, courseID, assignmentID int32) (*model.Assignment, error) {
	err := e.checkCourseFromTeacher(teacherID, courseID)
	if err != nil {
		return nil, err
	}
	assignment, err := e.assignmentsRepo.GetAssignment(assignmentID)
	if assignment == nil {
		return nil, fmt.Errorf("assignment with id %v does not exist", assignmentID)
	}
	if assignment.CourseID != courseID {
		return nil, fmt.Errorf("assignment with id %v does not belong to course with id %v", assignmentID, courseID)
	}
	return assignment, nil
}

func (e eUsecase) NewAssignment(teacherID, courseID int32, name string, comment *string) (*model.Assignment, error) {
	err := e.checkCourseFromTeacher(teacherID, courseID)
	if err != nil {
		return nil, err
	}
	assignment, err := e.assignmentsRepo.NewAssignment(courseID, name, comment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

func (e eUsecase) EditAssignment(
	teacherID,
	courseID,
	assignmentID int32,
	name string,
	comment *string,
	projectID *int32,
	finishedDate *time.Time,
	locked bool,
	answerConfig model.AssignmentAnswerConfig,
) error {
	_, err := e.GetAssignment(teacherID, courseID, assignmentID)
	if err != nil {
		return err
	}
	if projectID != nil {
		project, err := e.projectsRepo.GetProjectSuperficial(*projectID)
		if err != nil {
			return err
		}
		if project.OwnerID != nil && *project.OwnerID != teacherID {
			return fmt.Errorf(
				"the project with the id %v is not owned by you and cannot be used inside of this assignment",
				*projectID,
			)
		}
	}
	return e.assignmentsRepo.EditAssignment(assignmentID, name, comment, projectID, finishedDate, locked, answerConfig)
}

func (e eUsecase) EditAssignmentOrder(teacherID, courseID, assignmentID, newOrder int32) error {
	_, err := e.GetAssignment(teacherID, courseID, assignmentID)
	if err != nil {
		return err
	}
	amountOfAssignments, err := e.assignmentsRepo.GetAmountOfAssignmentsOfCourse(courseID)
	if err != nil {
		return err
	}
	// because order starts at 0 in a course, it should not be bigger than the amount of assignments in a course
	if newOrder >= amountOfAssignments {
		return fmt.Errorf(
			"the new order (%v) is invalid, "+
				"as the number of assignments in this course is %v and the order starts at 0",
			newOrder,
			amountOfAssignments,
		)
	}
	return e.assignmentsRepo.EditAssignmentOrder(courseID, assignmentID, newOrder)
}

func (e eUsecase) DeleteAssignment(teacherID, courseID, assignmentID int32) error {
	_, err := e.GetAssignment(teacherID, courseID, assignmentID)
	if err != nil {
		return err
	}
	return e.assignmentsRepo.DeleteAssignment(assignmentID)
}
