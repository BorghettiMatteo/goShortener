package routes

import (
	"context"
	. "main/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func ComputeShortener(c *fiber.Ctx) error {
	uri := c.Params("url")

	search := Database.HGet(ctx, uri, "expanded")
	if search == nil {
		return c.SendStatus(http.StatusNotFound)
	}
	c.SendString(search.Val())

	//redirect to unshortened url
	return c.Redirect(search.Val(), http.StatusPermanentRedirect)
}

func CreateShortener(c *fiber.Ctx) error {
	urlRequestBody := new(UrlRequestBody)
	if err := c.BodyParser(urlRequestBody); err != nil {
		return err
	}

	// save to redis
	//if item already in, return current item
	if dump := Database.Exists(ctx, urlRequestBody.Plain); dump.Val() != 0 {
		content := Database.HGetAll(ctx, urlRequestBody.Plain)
		c.SendStatus(http.StatusFound)
		return c.SendString(content.Val()["expanded"])
	}

	// create the entry on redis table
	encodedUrlValue := urlRequestBody.UrlEncoder()
	Database.HSet(ctx, encodedUrlValue, "expanded", urlRequestBody.Plain)
	c.SendStatus(http.StatusCreated)
	return c.SendString("0.0.0.0:8080/" + urlRequestBody.ApiVersion + "/shortener/" + encodedUrlValue)
}
