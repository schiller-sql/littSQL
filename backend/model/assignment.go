package model

type Assignment struct {
	//[ 0] project_id                                     INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ProjectID int32 `gorm:"primary_key;column:project_id;type:INT4;" json:"project_id"`
	//[ 1] course_id                                      INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	CourseID int32 `gorm:"primary_key;column:course_id;type:INT4;" json:"course_id"`
	//[ 2] status                                         USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
	Status string `gorm:"column:status;type:VARCHAR;" json:"status"`
	//[ 3] solution_mode                                  USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
	SolutionMode string `gorm:"column:solution_mode;type:VARCHAR;" json:"solution_mode"`
	//[ 4] sequence                                       INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Sequence int32 `gorm:"column:sequence;type:INT2;" json:"sequence"`
}
