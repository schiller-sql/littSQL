package courses

import (
	"github.com/schiller-sql/littSQL/model"
)

type Usecase interface {
	GetCoursesOfTeacher(teacherID int32) (*[]model.CourseListing, error)
	NewCourse(teacherID int32, name string) (*model.Course, error)
	GetCourseDetails(teacherID int32, courseID int32) (*model.Course, error)
	EditCourse(courseID int32, teacherID int32, name string) error
	DeleteCourse(teacherID int32, courseID int32) error
}
