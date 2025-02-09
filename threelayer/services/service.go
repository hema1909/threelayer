package services

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"threelayer/models"
	"threelayer/store"
)

type service struct {
	storeLayer store.Logger
}

func New(storeLayer store.Logger) service {
	return service{storeLayer: storeLayer}
}

func (s *service) Create(ctx *gofr.Context, logs models.Logs) (models.Logs, error) {
	logs.ID = uuid.New()

	supportedLevelValuesMap := map[string]bool{"HIGH": true, "LOW": true}
	_, ok := supportedLevelValuesMap[logs.Level]
	if !ok {
		return models.Logs{}, http.ErrorInvalidParam{Params: []string{"Level"}}
	}

	resp, err := s.storeLayer.Create(ctx, logs)
	if err != nil {
		return models.Logs{}, err
	}

	return resp, nil
}

func (s *service) GetByID(ctx *gofr.Context, ID uuid.UUID) (models.Logs, error) {

	resp, err := s.storeLayer.GetByID(ctx, ID)
	if err != nil {
		return models.Logs{}, err
	}

	return resp, nil
}
