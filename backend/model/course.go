package model

type Course struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] teacher_id                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TeacherID int32 `gorm:"column:teacher_id;type:INT4;" json:"teacher_id"`
	//[ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
}

type CourseListing struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
