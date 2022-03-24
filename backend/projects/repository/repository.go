package repository

import (
	"database/sql"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type eRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) projects.Repository {
	return &eRepository{DB: db}
}

func (e eRepository) GetProject(projectID int32, tasks bool) (*model.Project, error) {
	query := e.DB
	if tasks {
		// TODO: Preloading not working properly
		query = query.Preload(clause.Associations)
	}
	var project model.Project
	result := query.Find(&project, projectID)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &project, nil
}

func (e eRepository) GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error) {
	var projectsOfTeacher []model.ProjectListing
	result := e.DB.Raw(
		"select id, name, owner_id is null as is_public from projects where owner_id is null or owner_id = ?",
		teacherID,
	).Find(&projectsOfTeacher)
	return &projectsOfTeacher, result.Error
}

func (e eRepository) NewProject(teacherID int32, name string) (*model.Project, error) {
	project := model.Project{OwnerID: sql.NullInt64{Int64: int64(teacherID), Valid: true}, Name: name}
	result := e.DB.Create(&project)
	return &project, result.Error
}

func (e eRepository) SaveEditedProject(editedProject *model.Project) error {
	return e.DB.Select("*").Omit("id", "owner_id").Save(&editedProject).Error
}

func (e eRepository) DeleteProject(projectID int32) error {
	return e.DB.Delete(&model.Project{}, projectID).Error
}
