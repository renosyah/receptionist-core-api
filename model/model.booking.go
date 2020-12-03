package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	decimal "github.com/shopspring/decimal"
)

type (
	Booking struct {
		ID            uuid.UUID       `json:"id"`
		CustomerID    uuid.UUID       `json:"customer_id"`
		Total         decimal.Decimal `json:"total"`
		PaymentStatus int             `json:"payment_status"`
	}

	BookingResponse struct {
		ID            uuid.UUID       `json:"id"`
		CustomerID    uuid.UUID       `json:"customer_id"`
		Total         decimal.Decimal `json:"total"`
		PaymentStatus int             `json:"payment_status"`
	}
)

func (b *Booking) Response() BookingResponse {
	return BookingResponse{
		ID:            b.ID,
		CustomerID:    b.CustomerID,
		Total:         b.Total,
		PaymentStatus: b.PaymentStatus,
	}
}

func (b *Booking) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "booking" (customer_id,total,payment_status) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.CustomerID, b.Total, b.PaymentStatus).Scan(&b.ID)
	if err != nil {
		return b.ID, err
	}

	return b.ID, nil
}

func (b *Booking) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Booking, error) {
	all := []Booking{}

	query, args, err := param.Query(strings.Split("id,customer_id,total,payment_status", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,customer_id,total,payment_status FROM "booking" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Booking{}
		err = rows.Scan(
			&one.ID, &one.CustomerID, &one.Total, &one.PaymentStatus,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (b *Booking) One(ctx context.Context, db *sql.DB) (Booking, error) {
	one := Booking{}

	query := `SELECT id,customer_id,total,payment_status FROM "booking" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(
		&one.ID, &one.CustomerID, &one.Total, &one.PaymentStatus,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (b *Booking) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "booking" SET customer_id = $1,total = $2,payment_status = $3 WHERE id = $4 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.CustomerID, b.Total, b.PaymentStatus, b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (b *Booking) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "booking" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
