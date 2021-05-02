package transform

import (
	"github.com/yura-under-review/ports-domain-service/api"
	"github.com/yura-under-review/ports-domain-service/models"
)

func ToModelPort(src *api.PortInfo) *models.PortInfo {

	// TODO: implement
	return &models.PortInfo{}
}

func ToModelPorts(src []*api.PortInfo) []*models.PortInfo {

	res := make([]*models.PortInfo, 0, len(src))

	for _, port := range src {
		res = append(res, ToModelPort(port))
	}

	return res
}
