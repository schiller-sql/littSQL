package model

type Project struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 1] db_sql                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	DbSQL *string `gorm:"column:db_sql;type:INT4;" json:"db_sql"`
	//[ 2] name                                           VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Name string `gorm:"column:name;type:VARCHAR;" json:"name"`
	//[ 3] documentation_md                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	DocumentationMd string `gorm:"column:documentation_md;type:TEXT;" json:"documentation_md"`
	//[ 4] owner_id                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	OwnerID *int32 `gorm:"column:owner_id;type:INT4;" json:"owner_id"`

	Tasks []Task
}

type CachedProjectsSqlData struct {
	ProjectID int32  `gorm:"primary_key;AUTO_INCREMENT;column:project_id;type:INT4;" json:"id"`
	Data      []byte `gorm:"column:data;type:bytea;" json:"data"`
}

type ProjectListing struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	IsPublic bool   `json:"is_public"`
}
