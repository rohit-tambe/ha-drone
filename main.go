package main

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/ha-drone/app/config"

	"github.com/ha-drone/app/api"
	"github.com/labstack/echo"
)

func routes(serviceGroup *echo.Group) {
	serviceGroup.POST("/calculateLocation", api.CalculateLocation)
}
func main() {
	server := echo.New()
	// get sector id from env
	config.SectorID = 1
	serviceGroup := server.Group("/dns-service")
	routes(serviceGroup)
	server.Validator = &CustomValidator{validator: validator.New()}

	// Start the Server
	if err := server.Start(":8081"); err != nil {
		log.Fatal(err)
	}

}

// CustomValidator validate api request
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
