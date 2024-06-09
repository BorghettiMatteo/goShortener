package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePastaBin(c *fiber.Ctx) error {
	pastaBinRequest := new(models.PastaBin)

	dump := c.Body()
	json.Unmarshal(dump, &pastaBinRequest)

	//dump to db
	insertedId, err := pastaBinRequest.InsertPastaMexToDb()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if insertedId != "" {
		c.SendString("0.0.0.0:5555/" + insertedId)
	}
	return c.SendStatus(http.StatusCreated)
}

func GetPastaBin(c *fiber.Ctx) error {
	pastaBinIdValue := c.Params("pastabinid")
	objectID, err := primitive.ObjectIDFromHex(pastaBinIdValue)
	if err != nil {
		c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	var retrivedMex models.PastaBin
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = models.Db.Database("db").Collection("pastamexs").FindOne(ctx, filter).Decode(&retrivedMex)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.SendString(err.Error())
			return c.SendStatus(http.StatusInternalServerError)
		}
	}

	if redirect := c.Query("redirect"); redirect == "true" {
		c.Context().Redirect(retrivedMex.Body, http.StatusTemporaryRedirect)
	}
	c.JSON(retrivedMex.Body)
	return c.SendStatus(http.StatusFound)
}
