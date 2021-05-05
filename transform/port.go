package transform

import (
	"github.com/yura-under-review/ports-domain-service/api"
	"github.com/yura-under-review/ports-domain-service/models"
)

func ToModelPort(src *api.PortInfo) *models.PortInfo {

	var coord *models.Coordinate
	if src.Coordinate != nil {
		coord = &models.Coordinate{
			Latitude:  src.Coordinate.Latitude,
			Longitude: src.Coordinate.Longitude,
		}
	}

	return &models.PortInfo{
		Symbol:     src.Symbol,
		Name:       src.Name,
		City:       src.City,
		Province:   src.Province,
		Country:    src.Country,
		Alias:      src.Alias,
		Regions:    src.Regions,
		Timezones:  src.Timezones,
		Unlocks:    src.Unlocks,
		Code:       src.Code,
		Coordinate: coord,
	}
}

func ToModelPorts(src []*api.PortInfo) []*models.PortInfo {

	res := make([]*models.PortInfo, 0, len(src))

	for _, port := range src {
		res = append(res, ToModelPort(port))
	}

	return res
}
