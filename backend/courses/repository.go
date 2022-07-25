package courses

import "github.com/schiller-sql/littSQL/model"

type Repository interface {
	GetCoursesOfTeacher(teacherID int32) (*[]model.CourseListing, error)
	NewCourse(teacherID int32, name string) (*model.Course, error)
	GetCourse(courseID int32) (*model.Course, error)
	EditCourse(courseID int32, name string) error
	DeleteCourse(courseID int32) error
}
