package service

// MomCorpService organization
type MomCorpService struct {
	Value interface{} `json:"location"`
}

//GetLocation MomCorpService calculate logic
func (c MomCorpService) GetLocation(x, y, z, vel float64, sectorID int) interface{} {
	sID := float64(sectorID)
	return (x * sID) + (y * sID) + (z * sID) + (vel)
}

//GetResponse MomCorpService response logic
func (c MomCorpService) GetResponse(locationValue interface{}) interface{} {
	return MomCorpService{Value: locationValue}
}
