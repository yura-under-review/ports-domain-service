package repository

import (
	"context"

	"github.com/yura-under-review/ports-domain-service/models"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Init() error {

	return nil
}

func (r *Repository) PortUpsert(context.Context, *models.PortInfo) error {

	return nil
}

func (r *Repository) BatchPortUpsert(context.Context, []*models.PortInfo) error {
	return nil
}
