package api

import (
	"net/http"
	"strconv"

	"github.com/ha-drone/app/config"

	"github.com/ha-drone/app/service"

	"github.com/ha-drone/app/model"
	"github.com/labstack/echo"
)

// CalculateLocation calculate location controller
func CalculateLocation(c echo.Context) error {
	requestDTO := model.CalculateRequestDTO{}
	if err := c.Bind(&requestDTO); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	// validate
	if err := c.Validate(&requestDTO); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	// get Location service
	corp := service.Location{LocationService: service.AtlasCorp{}}

	// calculate location as per request coordinates
	locationVal := corp.GetLocation(GetFolatCoord(requestDTO.X), GetFolatCoord(requestDTO.Y),
		GetFolatCoord(requestDTO.Z), GetFolatCoord(requestDTO.Vel), config.SectorID)

	// get custome response
	response := corp.GetResponse(locationVal)
	return c.JSON(http.StatusOK, response)
}

// GetFolatCoord convert string into float
func GetFolatCoord(v string) float64 {
	if value, err := strconv.ParseFloat(v, 64); err == nil {
		return value
	}
	return 0.00
}
