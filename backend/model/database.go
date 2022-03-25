package model

type Database struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 2] data                                           BYTEA                null: false  primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []
	Data string `gorm:"column:data;type:BYTEA;" json:"data"`
	//[ 3] schema_svg_path                                VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	SchemaSVGPath *string `gorm:"column:schema_svg_path;type:VARCHAR;" json:"schema_svg_path"`
	//[ 4] owner_id                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	OwnerID *int32 `gorm:"column:owner_id;type:INT4;" json:"owner_id"`
}

type DatabaseListing struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
