package model

import uuid "github.com/satori/go.uuid"

type (
	CustomerAuth struct {
		SessionID        uuid.UUID `json:"session_id"`
		CustomerResponse `json:"customer"`
	}
)
