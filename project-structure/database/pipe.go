package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"pwa-workshop/project-structure/logger"
	"pwa-workshop/project-structure/model"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *Database) CreatePipe(ctx context.Context, req *model.Pipe) (uuid.UUID, error) {
	id := uuid.New()

	mgPipe := &model.MgPipe{
		Id:         id,
		Geometry:   req.Geometry,
		Properties: req.Properties,
		PgSynced:   false,
	}

	mgColl := d.MgCli.
		Database(mgDB).
		Collection(mgPipeCollection)

	_, err := mgColl.InsertOne(ctx, mgPipe)
	if err != nil {
		return uuid.Nil, fmt.Errorf("mongo insert failed: %w", err)
	}

	geoJSONBytes, err := json.Marshal(req.Geometry)
	if err != nil {
		return uuid.Nil, fmt.Errorf("geometry marshal failed: %w", err)
	}

	pgQuery := `
		INSERT INTO ` + pgPipeTable + ` (id, geom)
		VALUES ($1, ST_SetSRID(ST_GeomFromGeoJSON($2), 4326))
	`

	_, err = d.PgCli.ExecContext(ctx, pgQuery, id, string(geoJSONBytes))
	if err != nil {

		return uuid.Nil, fmt.Errorf("postgres insert failed: %w", err)
	}

	_, err = mgColl.UpdateByID(
		ctx,
		id,
		bson.M{
			"$set": bson.M{"pg_synced": true},
		},
	)
	if err != nil {
		logger.Warnf("warning: failed to update pg_synced for %s", id)
	}

	return id, nil
}

func (d *Database) GetPipeByID(ctx context.Context, id *uuid.UUID) (*model.MgPipe, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	col := d.MgCli.Database(mgDB).Collection(mgPipeCollection)

	var result model.MgPipe
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &result, nil
}

func (d *Database) ListPipes(
	ctx context.Context,
	filter bson.M,
) ([]model.MgPipe, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	col := d.MgCli.Database(mgDB).Collection(mgPipeCollection)

	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var pipes []model.MgPipe
	for cur.Next(ctx) {
		var p model.MgPipe
		if err := cur.Decode(&p); err != nil {
			return nil, err
		}
		pipes = append(pipes, p)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return pipes, nil
}

func (d *Database) UpdatePipeByID(
	ctx context.Context,
	id *uuid.UUID,
	update bson.M,
) error {

	if len(update) == 0 {
		return errors.New("update data is empty")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	col := d.MgCli.Database(mgDB).Collection(mgPipeCollection)

	res, err := col.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (d *Database) DeletePipeByID(
	ctx context.Context,
	id *uuid.UUID,
) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	col := d.MgCli.Database(mgDB).Collection(mgPipeCollection)

	res, err := col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
