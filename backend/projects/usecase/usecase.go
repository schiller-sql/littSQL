package usecase

import (
	"database/sql"
	"fmt"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
)

type usecase struct {
	repo projects.Repository
}

func NewUsecase(repo projects.Repository) projects.Usecase {
	return &usecase{repo}
}

func (u usecase) GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error) {
	project, err := u.repo.GetProject(projectID, true)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, fmt.Errorf("Project with the id '%d' could not be found", projectID)
	}
	if project.OwnerID.Valid && int32(project.OwnerID.Int64) != teacherID {
		return nil, fmt.Errorf("Not authorized to view the details of this project")
	}
	return project, err
}

func (u usecase) GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error) {
	return u.repo.GetProjectsOfTeacher(teacherID)
}

func (u usecase) NewProject(teacherID int32, name string) (*model.Project, error) {
	return u.repo.NewProject(teacherID, name)
}

func (u usecase) EditProject(
	projectID int32,
	teacherID int32,
	databaseID sql.NullInt64,
	name string,
	documentationMd string,
	tasks []model.Task,
) error {
	project, err := u.repo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("Project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if int32(project.OwnerID.Int64) != teacherID {
		return fmt.Errorf("Not authorized to delete this project")
	}
	return u.repo.SaveEditedProject(
		&model.Project{
			ID:              projectID,
			OwnerID:         sql.NullInt64{Int64: int64(teacherID), Valid: true},
			DatabaseID:      databaseID,
			Name:            name,
			DocumentationMd: documentationMd,
			Tasks:           tasks,
		},
	)
}

func (u usecase) DeleteProject(teacherID int32, projectID int32) error {
	project, err := u.repo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("Project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if int32(project.OwnerID.Int64) != teacherID {
		return fmt.Errorf("Not authorized to delete this project")
	}
	return u.repo.DeleteProject(projectID)
}
