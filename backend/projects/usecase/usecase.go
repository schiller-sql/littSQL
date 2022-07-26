package usecase

import (
	"fmt"

	model "github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/projects"
)

type eUsecase struct {
	projectsRepo projects.Repository
}

func NewUsecase(projectsRepo projects.Repository) projects.Usecase {
	return &eUsecase{projectsRepo}
}

func (u eUsecase) GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error) {
	return u.projectsRepo.GetProjectsOfTeacher(teacherID)
}

func (u eUsecase) NewProject(teacherID int32, name string) (*model.Project, error) {
	return u.projectsRepo.NewProject(teacherID, name)
}

func (u eUsecase) GetProjectDetails(teacherID int32, projectID int32) (*model.Project, error) {
	project, err := u.projectsRepo.GetProject(projectID)
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
	project, err := u.projectsRepo.GetProjectSuperficial(projectID)
	if project == nil {
		return fmt.Errorf("project with the id '%d' could not be found", projectID)
	}
	if err != nil {
		return err
	}
	if project.OwnerID == nil || *project.OwnerID != teacherID {
		return fmt.Errorf("not authorized to edit this project")
	}
	return u.projectsRepo.SaveEditedProject(
		&model.Project{
			ProjectSuperficial: model.ProjectSuperficial{
				ID:      projectID,
				OwnerID: &teacherID,
				Name:    name,
			},
			DbSQL:           sql,
			DocumentationMd: documentationMd,
			Tasks:           tasks,
		},
	)
}

func (u eUsecase) DeleteProject(teacherID int32, projectID int32) error {
	project, err := u.projectsRepo.GetProjectSuperficial(projectID)
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
