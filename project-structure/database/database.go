package database

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	MgCli *mongo.Client
	PgCli *sql.DB
}

func NewDatabase(MgCli *mongo.Client, PgCli *sql.DB) *Database {
	return &Database{
		MgCli: MgCli,
		PgCli: PgCli,
	}
}
