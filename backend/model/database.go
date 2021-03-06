package model

type DatabaseTemplate struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 2] name                                           TEXT null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Description string `gorm:"column:description;type:TEXT;" json:"description"`
	//[ 3] name                                           TEXT null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	SQL string `gorm:"column:sql;type:TEXT;" json:"sql"`
}

type DatabaseTemplateListing struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
