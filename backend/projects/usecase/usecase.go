package usecase

import (
	"fmt"
	sqlExecutor "github.com/schiller-sql/littSQL/sql_executor"

	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
)

type eUsecase struct {
	projectsRepo    projects.Repository
	sqlExecutorRepo sqlExecutor.Repository
}

func NewUsecase(projectsRepo projects.Repository, sqlExecutorRepo sqlExecutor.Repository) projects.Usecase {
	return &eUsecase{projectsRepo, sqlExecutorRepo}
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
		return nil, fmt.Errorf("project with the id '%d' could not be found", projectID)
	}
	if project.OwnerID != nil && *project.OwnerID != teacherID {
		return nil, fmt.Errorf("not authorized to view the details of this project")
	}
	return project, err
}

func (u eUsecase) EditProject(
	projectID int32,
	teacherID int32,
	name string,
	documentationMd string,
	sql *string,
	tasks []model.Task,
) error {
	project, err := u.projectsRepo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if project.OwnerID == nil || *project.OwnerID != teacherID {
		return fmt.Errorf("not authorized to edit this project")
	}
	err = u.projectsRepo.SaveEditedProject(
		&model.Project{
			ID:              projectID,
			OwnerID:         &teacherID,
			DbSQL:           sql,
			Name:            name,
			DocumentationMd: documentationMd,
			Tasks:           tasks,
		},
	)
	if err != nil {
		return err
	}
	var sqlData *[]byte
	if sql != nil {
		if sqlData, err = u.sqlExecutorRepo.ExecuteSQLite(*sql); err != nil {
			return err
		}
	}
	err = u.projectsRepo.SaveEditedProjectSqlData(projectID, sqlData)
	if err != nil {
		return err
	}
	return nil
}

func (u eUsecase) DeleteProject(teacherID int32, projectID int32) error {
	project, err := u.projectsRepo.GetProject(projectID, false)
	if project == nil {
		return fmt.Errorf("project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if project.OwnerID == nil || *project.OwnerID != teacherID {
		return fmt.Errorf("not authorized to delete this project")
	}
	return u.projectsRepo.DeleteProject(projectID)
}
