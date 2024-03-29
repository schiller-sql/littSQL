package usecase

import (
	"errors"
	"fmt"
	"github.com/schiller-sql/littSQL/courses"
	"github.com/schiller-sql/littSQL/model"
)

type eUsecase struct {
	coursesRepo courses.Repository
}

func NewUsecase(coursesRepo courses.Repository) courses.Usecase {
	return &eUsecase{coursesRepo}
}

func (u eUsecase) GetCoursesOfTeacher(teacherID int32) (*[]model.CourseListing, error) {
	return u.coursesRepo.GetCoursesOfTeacher(teacherID)
}

func (u eUsecase) NewCourse(teacherID int32, name string) (*model.Course, error) {
	if len(name) == 0 {
		return nil, errors.New("name is too short for a course")
	}
	return u.coursesRepo.NewCourse(teacherID, name)
}

func (u eUsecase) GetCourseDetails(teacherID int32, courseID int32) (*model.Course, error) {
	course, err := u.coursesRepo.GetCourse(courseID)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, fmt.Errorf("course with id '%d' not found", courseID)
	}
	if course.TeacherID != teacherID {
		return nil, errors.New("this course does not belong to your account")
	}
	return course, nil
}

func (u eUsecase) EditCourse(courseID int32, teacherID int32, name string) error {
	course, err := u.coursesRepo.GetCourse(courseID)
	if err != nil {
		return err
	}
	if course == nil {
		return fmt.Errorf("course with id '%d' not found", courseID)
	}
	if course.TeacherID != teacherID {
		return errors.New("this course does not belong to your account")
	}
	return u.coursesRepo.EditCourse(courseID, name)
}

func (u eUsecase) DeleteCourse(teacherID int32, courseID int32) error {
	course, err := u.coursesRepo.GetCourse(courseID)
	if err != nil {
		return err
	}
	if course == nil {
		return fmt.Errorf("course with id '%d' not found", courseID)
	}
	if course.TeacherID != teacherID {
		return errors.New("this course does not belong to your account")
	}
	return u.coursesRepo.DeleteCourse(courseID)
}
