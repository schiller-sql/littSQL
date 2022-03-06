package model

type Teacher struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 3] email                                          VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Email string `gorm:"column:email;type:VARCHAR;" json:"email"`
	//[ 4] password                                       VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Password string `gorm:"column:password;type:VARCHAR;" json:"password"`
}
