package model

import (
	"github.com/google/uuid"
)

// ใช้สำหรับ HTTP request
type Pipe struct {
	Id         string       `json:"id"`
	Geometry   GeoJSON      `json:"geometry"`
	Properties PipeProperty `json:"properties"`
}

type PipeProperty struct {
	Length float64 `json:"length"`
	Color  string  `json:"color"`
}

type GeoJSON struct {
	Type        string      `json:"type"`        // LineString
	Coordinates [][]float64 `json:"coordinates"` // [][]float64
}

// mongo db model
type MgPipe struct {
	Id         uuid.UUID    `bson:"_id"` // same logical ID as Pg
	Geometry   GeoJSON      `bson:"geometry"`
	Properties PipeProperty `bson:"properties"`
	PgSynced   bool         `bson:"pg_synced"`
}
