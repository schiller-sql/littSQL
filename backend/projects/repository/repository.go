package repository

import (
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

func (e eRepository) GetProjectsOfTeacher(teacherID int32) (*[]model.ProjectListing, error) {
	var projectsOfTeacher []model.ProjectListing
	result := e.DB.Raw(
		"select id, name, owner_id is null as is_public from projects where owner_id is null or owner_id = ?",
		teacherID,
	).Find(&projectsOfTeacher)
	return &projectsOfTeacher, result.Error
}

func (e eRepository) NewProject(teacherID int32, name string) (*model.Project, error) {
	project := model.Project{OwnerID: &teacherID, Name: name}
	result := e.DB.Create(&project)
	return &project, result.Error
}

func (e eRepository) GetProject(projectID int32, fillTasks bool) (*model.Project, error) {
	query := e.DB
	if fillTasks {
		query = query.Preload(clause.Associations)
	}
	var project model.Project
	result := query.Find(&project, projectID)
	err := result.Error
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	tasks := project.Tasks
	taskIds := make([]int32, len(tasks))
	for i := range taskIds {
		taskIds[i] = project.Tasks[i].Number
	}
	var questions []model.Question
	result = e.DB.
		Raw("select * from questions where project_id = ? and task_number in ?", projectID, taskIds).
		Scan(&questions)
	err = result.Error
	if err != nil {
		return nil, err
	}
	for _, question := range questions {
		for i, task := range tasks {
			if question.TaskNumber == task.Number {
				tasks[i].Questions = append(task.Questions, question)
				break
			}
		}
	}
	project.Tasks = tasks
	return &project, nil
}

func (e eRepository) SaveEditedProject(editedProject *model.Project) error {
	err := e.DB.Transaction(func(tx *gorm.DB) error {
		if err := e.DB.Where("project_id = ?", editedProject.ID).Delete(&model.Task{}).Error; err != nil {
			return err
		}
		if err := e.DB.Create(editedProject.Tasks).Error; err != nil {
			return err
		}
		if err := e.DB.Select("*").Omit("id", "owner_id").Save(&editedProject).Error; err != nil {
			return err
		}
		return nil
	})
	return err

}

func (e eRepository) DeleteProject(projectID int32) error {
	return e.DB.Delete(&model.Project{}, projectID).Error
}
