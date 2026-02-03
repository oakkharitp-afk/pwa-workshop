package migration

import (
	"context"
	"database/sql"
	"strings"
)

func MigratePipeSchema(ctx context.Context, db *sql.DB) error {
	var q = `

	CREATE SCHEMA IF NOT EXISTS <schema>;


	-- Enable PostGIS
	CREATE EXTENSION IF NOT EXISTS postgis;

	-- Pipe feature table (geometry)
	CREATE TABLE IF NOT EXISTS <schema>.<tableName> (
		id UUID PRIMARY KEY,
		geom geometry(LineString, 4326) NOT NULL
	);

	-- Spatial index
	CREATE INDEX IF NOT EXISTS idx_<tableName>_geom
	ON <schema>.<tableName>
	USING GIST (geom);
	`
	q = strings.ReplaceAll(q, "<tableName>", "pipe_geom")
	q = strings.ReplaceAll(q, "<schema>", "main")

	_, err := db.ExecContext(ctx, q)
	return err
}
