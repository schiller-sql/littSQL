package model

type Task struct {
	//[ 0] project_id                                     INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ProjectID int32 `gorm:"primary_key;column:project_id;type:INT4;" json:"project_id"`
	//[ 1] number                                         INT2                 null: false  primary: true   isArray: false  auto: false  col: INT2            len: -1      default: []
	Number int32 `gorm:"primary_key;column:number;type:INT2;" json:"number"`
	//[ 2] description                                    VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Description string `gorm:"column:description;type:VARCHAR;" json:"description"`
	//[ 3] is_voluntary                                   BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	IsVoluntary bool `gorm:"column:is_voluntary;type:BOOL;default:false;" json:"is_voluntary"`

	Questions []Question
}
