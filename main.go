package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"

	"github.com/quzhi1/fiber-versioning-tool/lib"
	"github.com/quzhi1/fiber-versioning-tool/schema"
	"github.com/quzhi1/fiber-versioning-tool/versions"
)

var versionChange = lib.VersionChangeList{
	versions.Version1_0,
	versions.Version1_1,
}

func main() {
	app := fiber.New()

	// Health
	app.Get("/", health)

	// Define endpoints
	app.Post(
		"/",
		versionChange.RequestVersionAdaptor,
		post,
		versionChange.ResponseVersionAdaptor,
	)

	// Start
	app.Listen(":8080")
}

func health(c *fiber.Ctx) error {
	_, err := c.WriteString("Healthy")
	return err
}

func post(c *fiber.Ctx) error {
	// Read request
	fmt.Printf("Received request: %v\n", string(c.Body()))
	var requestBody schema.RequestBodyVersion1_1

	// Parse request json
	err := json.Unmarshal(c.Body(), &requestBody)
	if err != nil {
		c.Status(400).WriteString("Invalid schema: " + err.Error())
		return err
	}

	// Construct response
	res := schema.ResponseBodyVersion1_1{
		Id:   utils.UUID(),
		Name: requestBody.FirstName + " " + requestBody.LastName,
	}

	// Return response
	fmt.Printf("Returning response: %v\n", res)
	err = c.JSON(res)
	if err != nil {
		c.Status(500).WriteString("Failed to construct JSON: " + err.Error())
		return err
	}

	// Set metadata response header
	c.Set("Client-Metadata", c.Get("Client-Metadata"))

	// Set region and lang response header
	if c.Query("region") != "" && c.Query("lang") != "" {
		c.Set("Language-Code", c.Query("lang")+"-"+c.Query("region"))
	}

	return c.Next()
}
