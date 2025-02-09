package services

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"threelayer/models"
)

type Logger interface {
	GetByID(ctx *gofr.Context, ID uuid.UUID) (models.Logs, error)
	Create(ctx *gofr.Context, request models.Logs) (models.Logs, error)
}
