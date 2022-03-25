package usecase

import (
	"database/sql"
	"fmt"
	"github.com/schiller-sql/littSQL/databases"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
)

type eUsecase struct {
	projectsRepo  projects.Repository
	databasesRepo databases.Repository
}

func NewUsecase(projectsRepo projects.Repository, databasesRepo databases.Repository) projects.Usecase {
	return &eUsecase{projectsRepo, databasesRepo}
}

func (u eUsecase) GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error) {
	return u.projectsRepo.GetProjectsOfTeacher(teacherID)
}

func (u eUsecase) NewProject(teacherID int32, name string) (*model.Project, error) {
	return u.projectsRepo.NewProject(teacherID, name)
}

func (u eUsecase) GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error) {
	project, err := u.projectsRepo.GetProject(projectID, true)
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

func (u eUsecase) EditProject(
	projectID int32,
	teacherID int32,
	databaseID sql.NullInt64,
	name string,
	documentationMd string,
	tasks []model.Task,
) error {
	project, err := u.projectsRepo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("Project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if int32(project.OwnerID.Int64) != teacherID {
		return fmt.Errorf("Not authorized to delete this project")
	}
	if databaseID.Valid {
		databaseID := int32(databaseID.Int64)
		database, err := u.databasesRepo.GetDatabase(databaseID, false)
		if err != nil {
			return err
		}
		if database == nil {
			return fmt.Errorf("The database with the id '%d' does not exist", databaseID)
		}
		if database.OwnerID != nil && *database.OwnerID != teacherID {
			return fmt.Errorf("The database with the id '%d' is not public", databaseID)
		}
	}
	return u.projectsRepo.SaveEditedProject(
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

func (u eUsecase) DeleteProject(teacherID int32, projectID int32) error {
	project, err := u.projectsRepo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("Project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if int32(project.OwnerID.Int64) != teacherID {
		return fmt.Errorf("Not authorized to delete this project")
	}
	return u.projectsRepo.DeleteProject(projectID)
}
