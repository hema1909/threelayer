package handler

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"threelayer/constants"
	"threelayer/models"
	"threelayer/services"
)

type handler struct {
	service services.Logger
}

func New(service services.Logger) handler {
	return handler{service: service}
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var logs models.Logs

	err := ctx.Bind(&logs)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Create(ctx, logs)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam(constants.ID)

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{[]string{"id"}}
	}
	resp, err := h.service.GetByID(ctx, parsedID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
