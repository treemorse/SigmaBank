package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jintonick/SigmaBank/database"
	"github.com/jintonick/SigmaBank/models"
	"strconv"
)

func CreateTasks(c *fiber.Ctx) error {
	var points []models.Point
	if err := database.DB.Find(&points).Error; err != nil {
		return err
	}

	for _, point := range points {
		if (point.Days > 7 && point.Approved > 0) || point.Days > 14 {
			_ = createTask("Выезд на точку для стимулирования выдач", point.Id, "Синьор", 4, 1)
		}
		if point.Cards > 0 && float64(point.Cards)/float64(point.Approved) < 0.5 {
			_ = createTask("Обучение агента", point.Id, "Мидл", 2, 2)
		}
		if !point.Materials {
			_ = createTask("Доставка карт и материалов", point.Id, "Любой", 1.5, 2)
		}
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func createTask(name string, pointID uint, grade string, duration float64, priority int) error {
	task := models.Task{
		Name:        name,
		PointID:     pointID,
		WorkerGrade: grade,
		Duration:    duration,
		Priority:    priority,
		WorkerID:    0,
	}
	err := CreateUniqueTask(database.DB, task)
	if err != nil {
		return err
	}
	return nil
}

func DistributeTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	if err := database.DB.Where("worker_id = 0").Order("priority asc, duration asc").Find(&tasks).Error; err != nil {
		return err
	}

	var workers []models.User
	if err := database.DB.Find(&workers).Error; err != nil {
		return err
	}

	for _, task := range tasks {
		var point models.Point
		if err := database.DB.Where("id = ?", task.PointID).First(&point).Error; err != nil {
			continue
		}

		bestWorker, err := findBestWorkerForTask(workers, task, point)
		if err != nil {
			continue
		}

		task.WorkerID = bestWorker.Id
		database.DB.Save(&task)

		bestWorker.AvailableHours -= task.Duration
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetUserTasks(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Query("userId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid user ID",
		})
	}

	var tasks []models.Task
	if err := database.DB.Where("worker_id = ?", userID).Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error fetching tasks",
		})
	}

	taskMaps := make([]map[string]string, 0)

	for _, task := range tasks {
		var point models.Point
		if err := database.DB.Where("id = ?", task.PointID).First(&point).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error fetching point",
			})
		}

		taskMap := map[string]string{
			"name":      task.Name,
			"longitude": strconv.FormatFloat(point.Longitude, 'f', 6, 64),
			"latitude":  strconv.FormatFloat(point.Latitude, 'f', 6, 64),
			"duration":  strconv.FormatFloat(task.Duration, 'f', 1, 64),
		}

		taskMaps = append(taskMaps, taskMap)
	}

	return c.JSON(taskMaps)
}
