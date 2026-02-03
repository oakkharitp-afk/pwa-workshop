package database

import (
	"context"
	"pwa-workshop/project-structure/model"
)

func (d *Database) CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (name, avatar)
		VALUES ($1, $2)
	`
	_, err := d.PgCli.ExecContext(
		ctx,
		query,
		user.Name,
		user.Avatar,
	)

	if err != nil {
		return err
	}

	return nil
}
