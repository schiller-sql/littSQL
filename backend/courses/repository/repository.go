package repository

import (
	"fmt"
	"github.com/schiller-sql/littSQL/courses"
	"github.com/schiller-sql/littSQL/model"
	"gorm.io/gorm"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) courses.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetCoursesOfTeacher(teacherID int32) (*[]model.CourseListing, error) {
	var coursesOfTeacher []model.CourseListing
	result := e.DB.Table("courses").Find(&coursesOfTeacher, &model.Course{TeacherID: teacherID})
	return &coursesOfTeacher, result.Error
}

func (e eRepository) NewCourse(teacherID int32, name string) (*model.Course, error) {
	course := model.Course{TeacherID: teacherID, Name: name}
	result := e.DB.Create(&course)
	return &course, result.Error
}

func (e eRepository) GetCourse(courseID int32) (*model.Course, error) {
	var course model.Course
	result := e.DB.Order("Upper(name)").Find(&course, courseID)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("course with id '%d' not found", courseID)
	}
	return &course, result.Error
}

func (e eRepository) EditCourse(courseID int32, name string) error {
	return e.DB.Select("*").Omit("teacher_id").Save(&model.Course{ID: courseID, Name: name}).Error
}

func (e eRepository) DeleteCourse(courseID int32) error {
	return e.DB.Delete(&model.Course{}, courseID).Error
}
