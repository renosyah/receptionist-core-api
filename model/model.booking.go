package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	decimal "github.com/shopspring/decimal"
)

type (
	Booking struct {
		ID            uuid.UUID       `json:"id"`
		CustomerID    uuid.UUID       `json:"customer_id"`
		SeatsID       uuid.UUID       `json:"seats_id"`
		Price         decimal.Decimal `json:"price"`
		Total         decimal.Decimal `json:"total"`
		DurationFrom  time.Time       `json:"duration_from"`
		DurationTo    time.Time       `json:"duration_to"`
		PaymentStatus int             `json:"payment_status"`
	}

	BookingResponse struct {
		ID            uuid.UUID       `json:"id"`
		CustomerID    uuid.UUID       `json:"customer_id"`
		SeatsID       uuid.UUID       `json:"seats_id"`
		Price         decimal.Decimal `json:"price"`
		Total         decimal.Decimal `json:"total"`
		DurationFrom  time.Time       `json:"duration_from"`
		DurationTo    time.Time       `json:"duration_to"`
		PaymentStatus int             `json:"payment_status"`
	}

	AllBooking struct {
		FilterBy    string `json:"filter_by"`
		FilterValue string `json:"filter_value"`
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (b *Booking) Response() BookingResponse {
	return BookingResponse{
		ID:            b.ID,
		CustomerID:    b.CustomerID,
		SeatsID:       b.SeatsID,
		Price:         b.Price,
		Total:         b.Total,
		DurationFrom:  b.DurationFrom,
		DurationTo:    b.DurationTo,
		PaymentStatus: b.PaymentStatus,
	}
}

func (b *Booking) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "booking" (customer_id,seats_id,price,total,duration_from,duration_to,payment_status) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query),
		b.CustomerID,
		b.SeatsID,
		b.Price,
		b.Total,
		b.DurationFrom,
		b.DurationTo,
		b.PaymentStatus,
	).Scan(&b.ID)
	if err != nil {
		return b.ID, err
	}

	return b.ID, nil
}

func (b *Booking) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Booking, error) {
	all := []Booking{}

	query, args, err := param.Query(strings.Split("id,customer_id,seats_id,price,total,duration_from,duration_to,payment_status", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,customer_id,seats_id,price,total,duration_from,duration_to,payment_status FROM "booking" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Booking{}
		err = rows.Scan(
			&one.ID, &one.CustomerID, &one.SeatsID, &one.Price, &one.Total, &one.DurationFrom, &one.DurationTo, &one.PaymentStatus,
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

	query := `SELECT id,customer_id,seats_id,price,total,duration_from,duration_to,payment_status FROM "booking" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(
		&one.ID, &one.CustomerID, &one.SeatsID, &one.Price, &one.Total, &one.DurationFrom, &one.DurationTo, &one.PaymentStatus,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (b *Booking) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "booking" SET customer_id = $1,seats_id = $2,price = $3,total = $4,duration_from = $5,duration_to = $6,payment_status = $7 WHERE id = $8 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.CustomerID, b.SeatsID, b.Price, b.Total, b.DurationFrom, b.DurationTo, b.PaymentStatus, b.ID).Scan(&id)
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
