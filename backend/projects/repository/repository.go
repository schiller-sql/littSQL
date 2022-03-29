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
		// gather important data to use later
		tasks := editedProject.Tasks
		taskNumbers := make([]int32, len(tasks))
		questions := make([]model.Question, 0, len(tasks)*2)
		for i, task := range tasks {
			taskNumbers[i] = task.Number
			for _, question := range task.Questions {
				questions = append(questions, question)
			}
		}
		// delete all redundant tasks
		if err := e.DB.Where("project_id = ? and number not in ?", editedProject.ID, taskNumbers).
			Delete(&model.Task{}).Error; err != nil {
			return err
		}
		// get all questions that are in the database
		var beforeQuestions []model.Question
		if err := e.DB.Where("project_id = ?", editedProject.ID).Find(&beforeQuestions).Error; err != nil {
			return err
		}
		// delete all redundant questions
		deltaQuestions := make([]model.Question, 0)
	bQ:
		for _, beforeQuestion := range beforeQuestions {
			for _, afterQuestion := range questions {
				if beforeQuestion.TaskNumber == afterQuestion.TaskNumber &&
					beforeQuestion.Number == afterQuestion.Number {
					continue bQ
				}
			}
			deltaQuestions = append(deltaQuestions, beforeQuestion)
		}
		for _, question := range deltaQuestions {
			if err := e.DB.Where(
				"project_id = ? and task_number = ? and number = ?",
				question.ProjectID,
				question.TaskNumber,
				question.Number,
			).Delete(&model.Question{}).Error; err != nil {
				return err
			}
		}
		// saving tasks, then questions and at last the project
		if err := e.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "project_id"}, {Name: "number"}},
			DoUpdates: clause.AssignmentColumns([]string{"description", "is_voluntary"}),
		}).Save(tasks).Error; err != nil {
			return err
		}
		if err := e.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "project_id"}, {Name: "task_number"}, {Name: "number"}},
			DoUpdates: clause.AssignmentColumns([]string{"question", "type", "solution"}),
		}).Save(questions).Error; err != nil {
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
