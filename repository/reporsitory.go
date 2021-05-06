package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yura-under-review/ports-domain-service/models"
)

type (
	Repository struct {
		config Config
		pool   *pgxpool.Pool
	}
)

const UpsertPortQuery = `insert into ports (symbol, name, country, province, city, alias, regions, timezones, unlocks, code, lat, lon)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
on conflict (symbol) do update set 
	name=excluded.name,
	province=excluded.province,
	city=excluded.city,
	alias=excluded.alias,
	regions=excluded.regions,
	timezones=excluded.timezones,
	unlocks=excluded.unlocks,
	code=excluded.code,
	lat=excluded.lat,
	lon=excluded.lon;
`

func New(config Config) *Repository {
	return &Repository{
		config: config,
	}
}

func (r *Repository) Init(ctx context.Context) error {

	pool, err := pgxpool.Connect(ctx, r.config.connectionString())
	if err != nil {
		return fmt.Errorf("failed to connect postgres: %w", err)
	}

	r.pool = pool

	return nil
}

func (r *Repository) Close() {
	r.pool.Close()
}

func (r *Repository) BatchPortUpsert(ctx context.Context, ports []*models.PortInfo) error {

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}

	defer conn.Release()

	batch := pgx.Batch{}

	for _, port := range ports {

		var lat, lon *float32
		if port.Coordinate != nil {
			lat = &port.Coordinate.Latitude
			lon = &port.Coordinate.Longitude
		}

		batch.Queue(UpsertPortQuery,
			port.Symbol,
			port.Name,
			port.Country,
			port.Province,
			port.City,
			sliceToString(port.Alias),
			sliceToString(port.Regions),
			sliceToString(port.Timezones),
			sliceToString(port.Unlocks),
			port.Code,
			lat,
			lon,
		)
	}

	batchResult := conn.SendBatch(ctx, &batch)
	defer func() { _ = batchResult.Close() }()

	if _, err := batchResult.Exec(); err != nil {

		return fmt.Errorf("failed to run query: %w", err)
	}

	return nil
}

func sliceToString(src []string) string {
	var res string
	for _, a := range src {
		res += a + ", "
	}

	if len(res) > 2 {
		res = res[:len(res)-2]
	}

	return res
}
