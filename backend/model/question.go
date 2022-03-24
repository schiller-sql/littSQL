package model

type Question struct {
	//[ 0] project_id                                     INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ProjectID int32 `gorm:"primary_key;column:project_id;type:INT4;" json:"project_id"`
	//[ 1] task_number                                    INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	TaskNumber int32 `gorm:"primary_key;column:task_number;type:INT4;" json:"task_number"`
	//[ 2] question                                       VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Question string `gorm:"column:question;type:VARCHAR;" json:"question"`
	//[ 3] number                                         INT2                 null: false  primary: true   isArray: false  auto: false  col: INT2            len: -1      default: []
	Number int32 `gorm:"primary_key;column:number;type:INT2;" json:"number"`
	//[ 4] type                                           USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
	Type string `gorm:"column:type;type:VARCHAR;" json:"type"`
	//[ 5] solution                                       VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Solution string `gorm:"column:solution;type:VARCHAR;" json:"solution"`
}
