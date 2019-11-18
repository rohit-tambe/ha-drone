package service

// AtlasCorp organization
type AtlasCorp struct {
	Value interface{} `json:"loc"`
}

//GetLocation AtlasCorp calculate logic
func (c AtlasCorp) GetLocation(x, y, z, vel float64, sectorID int) interface{} {
	sID := float64(sectorID)
	value := (x * sID) + (y * sID) + (z * sID) + (vel)
	return float32(value)
}

//GetResponse AtlasCorp response logic
func (c AtlasCorp) GetResponse(locationValue interface{}) interface{} {
	return AtlasCorp{Value: locationValue}
}
