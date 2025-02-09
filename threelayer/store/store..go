package store

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"threelayer/models"
)

type store struct {
	m map[string]models.Logs
}

func New(m map[string]models.Logs) store {
	return store{m: m}
}

func (s *store) Create(ctx *gofr.Context, logs models.Logs) (models.Logs, error) {
	s.m[logs.ID.String()] = logs

	return logs, nil
}

func (s *store) GetByID(ctx *gofr.Context, ID uuid.UUID) (models.Logs, error) {

	resp, ok := s.m[ID.String()]
	if !ok {
		return models.Logs{}, http.ErrorEntityNotFound{"ID", ID.String()}
	}

	return resp, nil

}
