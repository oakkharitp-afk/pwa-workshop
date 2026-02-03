package handler

import (
	"net/http"
	"pwa-workshop/project-structure/errs"
	"pwa-workshop/project-structure/logger"
	"pwa-workshop/project-structure/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) CreatePipe(c echo.Context) error {
	var pipe model.Pipe

	if err := c.Bind(&pipe); err != nil {
		logger.Error(err)
		return Error(c, errs.BadRequestError("invalid request body, ", err))
	}

	ctx := c.Request().Context()
	id, err := h.DB.CreatePipe(ctx, &pipe)
	if err != nil {
		return Error(c, err)
	}

	// อัปเดต ID ใน response
	pipe.Id = id.String()

	return c.JSON(http.StatusCreated, pipe)
}

func (h *Handler) GetPipes(c echo.Context) error {
	ctx := c.Request().Context()
	pipes, err := h.DB.ListPipes(ctx, bson.M{})
	if err != nil {
		return Error(c, err)
	}

	resp := make([]model.Pipe, len(pipes))

	for i, pipe := range pipes {
		resp[i] = model.Pipe{
			Id:         pipe.Id.String(),
			Geometry:   pipe.Geometry,
			Properties: pipe.Properties,
		}
	}
	return c.JSON(http.StatusOK, resp)
}
