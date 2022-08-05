package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"

	"github.com/quzhi1/fiber-versioning-tool/lib"
	"github.com/quzhi1/fiber-versioning-tool/schema"
	"github.com/quzhi1/fiber-versioning-tool/versions"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/quzhi1/fiber-versioning-tool/docs" // docs is generated by Swag CLI, you have to import it.
)

var versionChange = lib.VersionChangeList{
	versions.Version1_0,
	versions.Version1_1,
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// Health
	app.Get("/", health)

	// Handle swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Define endpoints
	app.Post(
		"/person",
		versionChange.RequestVersionAdaptor,
		post,
		versionChange.ResponseVersionAdaptor,
	)

	// Start
	app.Listen(":8080")
}

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func health(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"status": "healthy",
	})
}

// @Summary Post a name
// @Description Create a database entry with first and last name
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /person [post]
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
