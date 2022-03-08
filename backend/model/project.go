package model

import "database/sql"

type Project struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] database_id                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	DatabaseID sql.NullInt64 `gorm:"column:database_id;type:INT4;" json:"database_id"`
	//[ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 3] documentation_md                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DocumentationMd sql.NullString `gorm:"column:documentation_md;type:TEXT;" json:"documentation_md"`
	//[ 4] owner                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Owner sql.NullInt64 `gorm:"column:owner;type:INT4;" json:"owner"`
}