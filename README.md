HA DNS

# local setup
Clone into src/github.com folder

# code coverage
go test -v -coverprofile cover.out ./app/...
go tool cover -html=cover.out -o cover.html

# build docker image
docker build -t ha-drone .

# run docker
docker run -d -p 8081:8081 ha-drone

# public docker repository
pull image from rohit51288/ha-drone

# service implementation
type LocationService interface {
	GetLocation(x, y, z, vel float64, sectorID int) interface{}
	GetResponse(location interface{}) interface{}
}

func (l *Location) GetLocation(x, y, z, vel float64, sectorID int) interface{} {
	return l.LocationService.GetLocation(x, y, z, vel, sectorID)
}

func (l *Location) GetResponse(locationValue interface{}) interface{} {
	return l.LocationService.GetResponse(locationValue)
}
# load run time strategy
corp := service.Location{LocationService: service.AtlasCorp{}}

# corp defination
func (c AtlasCorp) GetLocation(x, y, z, vel float64, sectorID int) interface{} {
	sID := float64(sectorID)
	value := (x * sID) + (y * sID) + (z * sID) + (vel)
	return float32(value)
}

func (c AtlasCorp) GetResponse(locationValue interface{}) interface{} {
	return AtlasCorp{Value: locationValue}
}
