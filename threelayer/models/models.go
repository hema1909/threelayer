package models

import "github.com/google/uuid"

type Logs struct {
	ID         uuid.UUID `json:"ID"`
	Service    string    `json:"service"`
	Level      string    `json:"level"`
	LogMessage string    `json:"logMessage"`
}
