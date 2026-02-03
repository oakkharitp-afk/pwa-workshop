package tool

import (
	"errors"
	"pwa-workshop/project-structure/model"

	"github.com/google/uuid"
)

func ReqPipeToMgPipe(req *model.Pipe) (*model.MgPipe, error) {
	if req == nil {
		return nil, errors.New("pipe request is nil")
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid pipe id")
	}

	return &model.MgPipe{
		Id:         id,
		Geometry:   req.Geometry,
		Properties: req.Properties,
	}, nil
}

func MgPipeToReqPipe(mg *model.MgPipe) (*model.Pipe, error) {
	if mg == nil {
		return nil, errors.New("mongo pipe is nil")
	}

	return &model.Pipe{
		Id:         mg.Id.String(),
		Geometry:   mg.Geometry,
		Properties: mg.Properties,
	}, nil
}
