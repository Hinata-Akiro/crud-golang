package tasks

import (
    "time"
	"crud-app/database"
)


func createNewTask(input *Task) (*Task, error) {
	now := time.Now().UTC()

	task := &Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		CreatedAt:   now,
		UpdatedAt:   now,
		DueDate:     input.DueDate,
		Assignee:    input.Assignee,
	}

	if err := database.Db.Create(&task).Error; err != nil {
		return nil ,err
	}
	return task, nil
}

func allTasks() ([]Task, error) {
	var tasks []Task

    if err := database.Db.Find(&tasks).Error; err != nil {
        return nil ,err
    }
    return tasks, nil
}

func getSingleTask(taskId uint64) (*Task, error) {
	var task Task
    if err := database.Db.Where("id = ?", taskId).First(&task).Error; err != nil {
        return nil ,err
    }
    return &task, nil
}

func editTaskById(taskId uint64, input *Task) (*Task, error) {
    var existingTask Task
    if err := database.Db.Where("id = ?", taskId).First(&existingTask).Error; err != nil {
        return nil, err
    }

    if input != nil {
        if input.Title != "" {
            existingTask.Title = input.Title
        }
        if input.Description != "" {
            existingTask.Description = input.Description
        }
        if input.Status != "" {
            existingTask.Status = input.Status
        }
        if input.Priority != "" {
            existingTask.Priority = input.Priority
        }
        if !input.DueDate.IsZero() {
            existingTask.DueDate = input.DueDate
        }
        if input.Assignee != "" {
            existingTask.Assignee = input.Assignee
        }
    }

    if err := database.Db.Save(&existingTask).Error; err != nil {
        return nil, err
    }

    return &existingTask, nil
}

func deleteTask(taskId uint64) error {
	var task Task
    if err := database.Db.Where("id = ?", taskId).First(&task).Error; err != nil {
        return err
    }

    if err := database.Db.Delete(&task).Error; err != nil {
        return err
    }
    return nil
}