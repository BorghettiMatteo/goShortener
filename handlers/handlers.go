package handlers

import (
	"encoding/json"
	"net/http"

	"main/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePastaBin(c *fiber.Ctx) error {
	pastaBinRequest := new(models.PastaBin)

	dump := c.Body()
	json.Unmarshal(dump, &pastaBinRequest)

	//dump to db
	err, insertedId := pastaBinRequest.InsertPastaMexToDb()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if insertedId != "" {
		c.SendString("0.0.0.0:5555/" + insertedId)
	}
	return c.SendStatus(http.StatusCreated)
}
