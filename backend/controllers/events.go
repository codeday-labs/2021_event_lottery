package controllers

import (
	"fmt"
	"github.com/codeday-labs/2021_event_lottery/database"
	"github.com/codeday-labs/2021_event_lottery/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	database.Connection.Find(&events)
	return c.JSON(events)
}

func GetEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event
	database.Connection.Find(&event, id)
	return c.JSON(event)
}

func CreateEvent(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Parsing Body
	maxAttendees, err := strconv.Atoi(data["maxAttendees"])
	if err != nil {
		fmt.Println(err)
	}

	event := models.Event{
		EventName:    data["eventName"],
		MaxAttendees: maxAttendees,
		StartDate:    data["startDate"],
		StartTime:    data["startTime"],
		EndDate:      data["endDate"],
		EndTime:      data["endTime"],
		LotteryDate:  data["lotteryDate"],
		LotteryTime:  data["lotteryTime"],
	}

	database.Connection.Create(&event)

	return c.JSON(event)
}