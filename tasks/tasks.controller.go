package tasks

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"crud-app/utils"
)





func TaskController (app * fiber.App) {
	taskRoute := app.Group("/tasks")
    taskRoute.Post("/", newTask)
	taskRoute.Get("/", getAllTasks)
	taskRoute.Get("/:taskId",  getTaskById)
	taskRoute.Put("/:taskId", updateTaskById)
	taskRoute.Delete("/:taskId", deleteTaskById)
}



func newTask(c *fiber.Ctx) error {
    var input Task
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    if err := input.Validate(); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Validation failed",
            Status:  fiber.StatusBadRequest,
        })
    }

    task, err := createNewTask(&input)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Error creating task",
            Status:  fiber.StatusBadRequest,
        })
    }

    return c.Status(fiber.StatusCreated).JSON(utils.Response{
        Message: "Task created successfully",
        Data:    &fiber.Map{"task": task},
        Status:  fiber.StatusCreated,
    })
}

func getTaskById(c *fiber.Ctx) error {
    taskIdStr := c.Params("taskId")

    taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

    task, err := getSingleTask(taskId)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(utils.Response{
            Message: "Task not found",
            Status:  fiber.StatusNotFound,
        })
    }

    return c.Status(fiber.StatusOK).JSON(utils.Response{
        Message: "Task found",
        Data:    &fiber.Map{"task": task},
        Status:  fiber.StatusOK,
    })
}

func getAllTasks(c *fiber.Ctx) error {
    tasks, err := allTasks()
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(utils.Response{
            Message: "Tasks not found",
            Status:  fiber.StatusNotFound,
        })
    }
    return c.Status(fiber.StatusOK).JSON(utils.Response{
        Message: "Tasks found",
        Data:    &fiber.Map{"tasks": tasks},
        Status:  fiber.StatusOK,
    })
}

func updateTaskById(c *fiber.Ctx) error {
    var input Task
    taskIdStr := c.Params("taskId")

    taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    existingTask, editErr := editTaskById(taskId, &input)
    if editErr != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Error updating task",
            Status:  fiber.StatusBadRequest,
        })
    }

    return c.Status(fiber.StatusOK).JSON(utils.Response{
        Message: "Task updated successfully",
        Data:    &fiber.Map{"task": existingTask},
        Status:  fiber.StatusOK,
    })
}

func deleteTaskById(c *fiber.Ctx) error {
    taskIdStr := c.Params("taskId")

    taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
    }

    deleteErr := deleteTask(taskId)
    if deleteErr != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Error deleting task",
            Status:  fiber.StatusBadRequest,
        })
    }

    return c.Status(fiber.StatusNoContent).SendString("Task deleted successfully")
}