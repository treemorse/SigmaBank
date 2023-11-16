package controllers

import (
	"errors"
	"fmt"
	"github.com/jintonick/SigmaBank/models"
	"gorm.io/gorm"
	"math"
	"sort"
	"strconv"
	"strings"
)

func convertToFloat64Slice(input string) ([]float64, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("input string must have exactly two parts")
	}

	var coords []float64
	for _, part := range parts {
		val, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}
		coords = append(coords, val)
	}

	return coords, nil
}

func CreateOrUpdatePoint(db *gorm.DB, newPoint models.Point) error {
	var existingPoint models.Point
	if err := db.Where("longitude = ? AND latitude = ?", newPoint.Longitude, newPoint.Latitude).First(&existingPoint).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return db.Create(&newPoint).Error
		}
		return err
	}

	return db.Model(&existingPoint).Updates(newPoint).Error
}

func CreateUniqueTask(db *gorm.DB, newTask models.Task) error {
	var existingTask models.Task
	if err := db.Where("name = ? AND point_id = ? AND worker_id = ?", newTask.Name, newTask.PointID, newTask.WorkerID).First(&existingTask).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return db.Create(&newTask).Error
		}
		return err
	}
	return nil
}

func findBestWorkerForTask(workers []models.User, task models.Task, point models.Point) (*models.User, error) {
	sort.Slice(workers, func(i, j int) bool {
		return distance(workers[i].Latitude, workers[i].Longitude, point.Latitude, point.Longitude) <
			distance(workers[j].Latitude, workers[j].Longitude, point.Latitude, point.Longitude)
	})

	for _, worker := range workers {
		if worker.Grade == task.WorkerGrade && worker.AvailableHours >= task.Duration {
			return &worker, nil
		}
	}

	return nil, errors.New("no suitable worker found")
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	return math.Sqrt(math.Pow(lat2-lat1, 2) + math.Pow(lon2-lon1, 2))
}
