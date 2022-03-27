package model

type Participant struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] course_id                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CourseID int32 `gorm:"column:course_id;type:INT4;" json:"course_id"`
	//[ 2] name                                           VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name *string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 3] access_code                                    BPCHAR(6)            null: false  primary: false  isArray: false  auto: false  col: BPCHAR          len: 6       default: [utils.random_string(6)]
	AccessCode string `gorm:"column:access_code;type:BPCHAR;size:6;default:utils.random_string(6);" json:"access_code"`
}
