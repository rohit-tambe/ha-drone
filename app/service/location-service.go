package service

import (
	"errors"
	"fmt"
)

// LocationService service defination
type LocationService interface {
	GetLocation(x, y, z, vel float64, sectorID int) interface{}
	GetResponse(location interface{}) interface{}
}

// Location location service
type Location struct {
	LocationService LocationService
}

// GetLocation calculate formula
func (l *Location) GetLocation(x, y, z, vel float64, sectorID int) interface{} {
	return l.LocationService.GetLocation(x, y, z, vel, sectorID)
}

// GetResponse custome response
func (l *Location) GetResponse(locationValue interface{}) interface{} {
	return l.LocationService.GetResponse(locationValue)
}

// GetCorp get organization
// for future enhancement
func GetCorp(corpID string) (Location, error) {
	switch corpID {
	case "ATLAS":
		return Location{AtlasCorp{}}, nil
	case "MomCorp":
		return Location{MomCorpService{}}, nil
	default:
		exception := fmt.Sprintf("Location service %v not implemented", corpID)
		return Location{}, errors.New(exception)
	}
}
