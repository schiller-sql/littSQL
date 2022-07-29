package repository

import (
	"github.com/schiller-sql/littSQL/assignments"
	"github.com/schiller-sql/littSQL/model"
	"gorm.io/gorm"
	"time"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) assignments.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetAssignmentsOfCourse(courseID int32) (*[]model.AssignmentListing, error) {
	var assignmentsOfCourse []model.AssignmentListing
	err := e.DB.Table("assignments_listing").Order("number").Find(&assignmentsOfCourse, &model.Assignment{CourseID: courseID}).Error
	if err != nil {
		return nil, err
	}
	return &assignmentsOfCourse, nil
}

func (e eRepository) GetAmountOfAssignmentsOfCourse(courseID int32) (int32, error) {
	var amountOfAssignments int64
	err := e.DB.Model(&model.Assignment{}).
		Where("assignments.course_id = ?", courseID).
		Count(&amountOfAssignments).Error
	if err != nil {
		return 0, err
	}
	return int32(amountOfAssignments), nil
}

func (e eRepository) NewAssignment(courseID int32, name string, ) (*model.Assignment, error) {
	amount, err := e.GetAmountOfAssignmentsOfCourse(courseID)
	if err != nil {
		return nil, err
	}
	assignment := model.Assignment{Name: name, CourseID: courseID, Number: amount}
	err = e.DB.Create(&assignment).Error
	return &assignment, err
}

func (e eRepository) GetAssignment(assignmentID int32) (*model.Assignment, error) {
	var assignment model.Assignment
	result := e.DB.Find(&assignment, assignmentID)
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &assignment, nil
}

func (e eRepository) EditAssignment(
	assignmentID int32,
	name string,
	comment *string,
	projectID *int32,
	finishedDate *time.Time,
	locked bool,
	answerConfig model.AssignmentAnswerConfig,
) error {
	return e.DB.Omit("course_id, number").Save(
		&model.Assignment{
			ID:                     assignmentID,
			Name:                   name,
			Comment:                comment,
			ProjectID:              projectID,
			FinishedDate:           finishedDate,
			Locked:                 locked,
			AnswerConfig: answerConfig,
		},
	).Error
}

func (e eRepository) EditAssignmentOrder(courseID int32, assignmentID int32, newOrder int32) error {
	return e.DB.Transaction(func(tx *gorm.DB) error {
		assignment := model.Assignment{ID: assignmentID}
		if err := tx.First(&assignment).Error; err != nil {
			return err
		}
		oldOrder := assignment.Number
		var moveUp bool
		if oldOrder > newOrder {
			moveUp = true
		} else {
			moveUp = false
		}
		var err error
		if moveUp {
			err = tx.
				Exec(
					"update assignments set number = number + 1 where course_id = ? and number >= ? and number < ?",
					courseID,
					newOrder,
					oldOrder,
				).Error
		} else {
			err = tx.
				Exec(
					"update assignments set number = number - 1 where course_id = ? and number > ? and number <= ?",
					courseID,
					oldOrder,
					newOrder,
				).Error
		}
		if err != nil {
			return err
		}
		assignment.Number = newOrder
		return tx.Select("number").Save(assignment).Error
	})
}

func (e eRepository) DeleteAssignment(assignmentID int32) error {
	return e.DB.Delete(&model.Assignment{}, assignmentID).Error
}
