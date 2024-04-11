package tasks

import (
	"errors"
	"time"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

)

type Task struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
	Assignee    string    `json:"assignee"`
}

var input struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	Assignee    string    `json:"assignee"`
}

func validateTaskInputs(title, description, status, priority string, dueDate time.Time, assignee string) error {

	if title == "" {
		return errors.New("title is required")
	}

	if description == "" {
		return errors.New("description is required")
	}

	if status == "" {
		return errors.New("status is required")
	}

	if priority == "" {
		return errors.New("priority is required")
	}

	if dueDate.IsZero() {
		return errors.New("due date is required")
	}

	if assignee == "" {
		return errors.New("assignee is required")
	}

	return nil
}

func NewTask(c *fiber.Ctx , db *gorm.DB)  error {
	if err := c.BodyParser(&input); err!= nil {
        return c.Status(400).SendString(err.Error())
    }
	if err := validateTaskInputs(input.Title, input.Description, input.Status, input.Priority, input.DueDate, input.Assignee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


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

	if err := db.Create(&task).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}


func GetTaskById(c *fiber.Ctx, db *gorm.DB) error {
	taskIdStr := c.Params("taskId")

	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)

	if err!= nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

	var task Task
    if err := db.Where("id = ?", taskId).First(&task).Error; err != nil {
        return c.Status(fiber.StatusNotFound).SendString("Task not found")
    }

    return c.Status(fiber.StatusOK).JSON(task)

}

func GetAllTasks(c *fiber.Ctx, db *gorm.DB) error {
	var tasks []Task
    if err := db.Find(&tasks).Error; err!= nil {
        return err
    }
    return c.Status(fiber.StatusFound).JSON(tasks)
}

func UpdateTaskById (c * fiber.Ctx, db *gorm.DB) error {
	taskIdStr := c.Params("taskId")

    taskId, err := strconv.ParseUint(taskIdStr, 10, 64)

    if err!= nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

    
    if err := c.BodyParser(&input); err!= nil {
        return c.Status(400).SendString(err.Error())
    }

	var existingTask Task
    if err := db.Where("id = ?", taskId).First(&existingTask).Error; err != nil {
        return c.Status(fiber.StatusNotFound).SendString("Task not found")
    }

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

    if err := db.Save(&existingTask).Error; err != nil {
        return err
    }

    return c.Status(fiber.StatusOK).JSON(existingTask)
}

func DeleteTaskById(c * fiber.Ctx, db *gorm.DB) error {
	taskIdStr := c.Params("taskId")

    taskId, err := strconv.ParseUint(taskIdStr, 10, 64)

    if err!= nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

    var existingTask Task
    if err := db.Where("id = ?", taskId).First(&existingTask).Error; err != nil {
        return c.Status(fiber.StatusNotFound).SendString("Task not found")
    }

    if err := db.Delete(&existingTask).Error; err != nil {
        return err
    }

    return c.Status(fiber.StatusNoContent).JSON(existingTask)
}