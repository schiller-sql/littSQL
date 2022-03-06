package model

import "time"

type Answer struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] course_id                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CourseID int32 `gorm:"column:course_id;type:INT4;" json:"course_id"`
	//[ 2] participant_id                                 INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	ParticipantID int32 `gorm:"column:participant_id;type:INT4;" json:"participant_id"`
	//[ 3] project_id                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	ProjectID int32 `gorm:"column:project_id;type:INT4;" json:"project_id"`
	//[ 4] task_number                                    INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	TaskNumber int32 `gorm:"column:task_number;type:INT2;" json:"task_number"`
	//[ 5] question_number                                INT2                 null: false  primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	QuestionNumber int32 `gorm:"column:question_number;type:INT2;" json:"question_number"`
	//[ 6] created_at                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [now()]
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//[ 7] answer                                         VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Answer string `gorm:"column:answer;type:VARCHAR;" json:"answer"`
	//[ 8] is_correct_automatic                           USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: [unknown]
	IsCorrectAutomatic string `gorm:"column:is_correct_automatic;type:VARCHAR;default:unknown;" json:"is_correct_automatic"`
	//[ 9] is_correct_manual_approval                     USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: [unknown]
	IsCorrectManualApproval string `gorm:"column:is_correct_manual_approval;type:VARCHAR;default:unknown;" json:"is_correct_manual_approval"`
}
