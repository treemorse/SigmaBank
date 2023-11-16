package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jintonick/SigmaBank/database"
	"github.com/jintonick/SigmaBank/models"
	"strconv"
)

func CreatePoint(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	coords, err := convertToFloat64Slice(data["coordinates"])
	if err != nil {
		return err
	}

	days, err := strconv.ParseUint(data["days"], 10, 32)
	if err != nil {
		return err
	}

	approved, err := strconv.ParseUint(data["approved"], 10, 32)
	if err != nil {
		return err
	}

	cards, err := strconv.ParseUint(data["cards"], 10, 32)
	if err != nil {
		return err
	}

	activated := data["activated"] == "вчера"
	materials := data["materials"] == "да"

	point := models.Point{
		Longitude: coords[0],
		Latitude:  coords[1],
		Activated: activated,
		Materials: materials,
		Days:      uint32(days),
		Approved:  uint32(approved),
		Cards:     uint32(cards),
	}

	err = CreateOrUpdatePoint(database.DB, point)
	if err != nil {
		return err
	}

	return c.JSON(point)
}
