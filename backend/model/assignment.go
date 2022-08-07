package model

import (
	"time"
)

// Assignment struct is a row record of the assignments table in the postgres database
type Assignment struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] number                                         INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Number int32 `gorm:"column:number;type:INT2;" json:"-"`
	//[ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 3] comment                                        VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Comment *string `gorm:"column:comment;type:VARCHAR;" json:"comment"`
	//[ 4] course_id                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CourseID int32 `gorm:"column:course_id;type:INT4;" json:"-"`
	//[ 5] project_id                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	ProjectID *int32 `gorm:"column:project_id;type:INT4;" json:"project_id"`
	//[ 6] finished_date                                  TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	FinishedDate *time.Time `gorm:"column:finished_date;type:TIMESTAMP;" json:"finished_date"`
	//[13] locked                                         BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	Locked bool `gorm:"column:locked;type:BOOL;default:true;" json:"locked"`

	AnswerConfig AssignmentAnswerConfig `gorm:"embedded" json:"answer_config"`
}
type AssignmentAnswerConfig struct {
	//[ 7] enable_auto_correction_on_sql_questions        BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	EnableAutoCorrectionOnSqlQuestions bool `gorm:"column:enable_auto_correction_on_sql_questions;type:BOOL;default:true;" json:"enable_auto_correction_on_sql_questions"`
	//[ 8] show_query_solution                            BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	ShowQuerySolution bool `gorm:"column:show_query_solution;type:BOOL;default:true;" json:"show_query_solution"`
	//[ 9] submit_only_once                               BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [true]
	SubmitOnlyOnce bool `gorm:"column:submit_only_once;type:BOOL;default:true;" json:"submit_only_once"`
	//[10] active_correction_behaviour                    USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: [show_correction]
	ActiveCorrectionBehaviour string `gorm:"column:active_correction_behaviour;type:VARCHAR;default:show_correction;" json:"active_correction_behaviour"`
	//[11] finished_correction_behavior                   USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: [show_correction_and_solution]
	FinishedCorrectionBehavior string `gorm:"column:finished_correction_behavior;type:VARCHAR;default:show_correction_and_solution;" json:"finished_correction_behavior"`
	//[12] finished_hide_answers                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	FinishedHideAnswers bool `gorm:"column:finished_hide_answers;type:BOOL;default:false;" json:"finished_hide_answers"`
}