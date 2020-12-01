package model

import (
	uuid "github.com/satori/go.uuid"
	mid "github.com/veritrans/go-midtrans"
)

type (
	Midtrans struct {
		BookingID          uuid.UUID              `json:"booking_id"`
		CustomerID         uuid.UUID              `json:"customer_id"`
		TransactionDetails mid.TransactionDetails `json:"transaction"`
		Items              []mid.ItemDetail       `json:"items"`
	}
)
