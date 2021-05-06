package grpc

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/yura-under-review/ports-domain-service/api"
	"github.com/yura-under-review/ports-domain-service/models"
	"github.com/yura-under-review/ports-domain-service/transform"
)

type PortsRepository interface {
	BatchPortUpsert(context.Context, []*models.PortInfo) error
}

type Resolver struct {
	api.UnimplementedPortsDomainServiceServer

	repo PortsRepository
}

func NewResolver(repo PortsRepository) *Resolver {
	return &Resolver{
		repo: repo,
	}
}

func (r *Resolver) UpsertPort(ctx context.Context, req *api.PortInfoRequest) (*api.PortInfoResponse, error) {
	modelPort := transform.ToModelPort(req.Port)

	if err := r.repo.BatchPortUpsert(ctx, []*models.PortInfo{modelPort}); err != nil {
		log.Errorf("faield ot upsert ports: %v", err)
		return nil, errors.New("failed to upsert ports")
	}

	return &api.PortInfoResponse{}, nil
}

func (r *Resolver) BatchUpsertPorts(ctx context.Context, req *api.BatchUpsertPortsRequest) (*api.BatchUpsertPortsResponse, error) {

	modelPorts := transform.ToModelPorts(req.Ports)

	if err := r.repo.BatchPortUpsert(ctx, modelPorts); err != nil {
		log.Errorf("faield ot upsert ports: %v", err)
		return nil, errors.New("failed to upsert ports")
	}

	log.Debugf("request was successfully processed [Nports: %d]", len(req.Ports))

	return &api.BatchUpsertPortsResponse{
		FailedPorts: nil,
	}, nil
}
