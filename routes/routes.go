package routes

import (
	"context"
	. "main/models"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func ComputeShortener(c *fiber.Ctx) error {
	param := c.Params("url")
	return c.SendString(param)
}

func CreateShortener(c *fiber.Ctx) error {
	urlRequest := new(UrlRequest)
	if err := c.BodyParser(urlRequest); err != nil {
		return err
	}

	// save to redis
	//if item already in, return current item
	if dump := Database.Exists(ctx, urlRequest.Plain); dump.Val() != 0 {
		content := Database.HGetAll(ctx, urlRequest.Plain)
		c.SendStatus(http.StatusFound)
		return c.SendString(content.Val()["shortened"])
	}

	// create the entry on redis table
	urlRequest.UrlEncoder(os.Getenv("CURRENT_API_VERSION"))
	Database.HSet(ctx, urlRequest.Plain, "shortened", urlRequest.Shortened)
	return c.SendStatus(http.StatusCreated)
}
