package handler

import "pwa-workshop/project-structure/database"

type Handler struct {
	DB *database.Database
}

func NewHandler(DB *database.Database) *Handler {
	return &Handler{DB: DB}
}
