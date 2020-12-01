package model

import uuid "github.com/satori/go.uuid"

type (
	OwnerAuth struct {
		SessionID     uuid.UUID `json:"session_id"`
		OwnerResponse `json:"owner"`
	}
)
