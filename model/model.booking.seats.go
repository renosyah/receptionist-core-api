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
	BookingSeats struct {
		ID           uuid.UUID       `json:"id"`
		BookingID    uuid.UUID       `json:"booking_id"`
		SeatsID      uuid.UUID       `json:"seats_id"`
		Price        decimal.Decimal `json:"price"`
		SubTotal     decimal.Decimal `json:"sub_total"`
		DurationFrom time.Time       `json:"duration_from"`
		DurationTo   time.Time       `json:"duration_to"`
	}

	BookingSeatsResponse struct {
		ID           uuid.UUID       `json:"id"`
		BookingID    uuid.UUID       `json:"booking_id"`
		SeatsID      uuid.UUID       `json:"seats_id"`
		Price        decimal.Decimal `json:"price"`
		SubTotal     decimal.Decimal `json:"sub_total"`
		DurationFrom time.Time       `json:"duration_from"`
		DurationTo   time.Time       `json:"duration_to"`
	}
)

func (b *BookingSeats) Response() BookingSeatsResponse {
	return BookingSeatsResponse{
		ID:           b.ID,
		BookingID:    b.BookingID,
		SeatsID:      b.SeatsID,
		Price:        b.Price,
		SubTotal:     b.SubTotal,
		DurationFrom: b.DurationFrom,
		DurationTo:   b.DurationTo,
	}
}

func (b *BookingSeats) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "booking_seats" (booking_id,seats_id,price,sub_total,duration_from,duration_to) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.BookingID, b.SeatsID, b.Price, b.SubTotal, b.DurationFrom, b.DurationTo).Scan(&b.ID)
	if err != nil {
		return b.ID, err
	}

	return b.ID, nil
}

func (b *BookingSeats) All(ctx context.Context, db *sql.DB, param ListQuery) ([]BookingSeats, error) {
	all := []BookingSeats{}

	query, args, err := param.Query(strings.Split("id,booking_id,seats_id,price,sub_total,duration_from,duration_to", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,booking_id,seats_id,price,sub_total,duration_from,duration_to FROM "booking_seats" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := BookingSeats{}
		err = rows.Scan(
			&one.ID, &one.BookingID, &one.SeatsID, &one.Price, &one.SubTotal, &one.DurationFrom, &one.DurationTo,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (b *BookingSeats) One(ctx context.Context, db *sql.DB) (BookingSeats, error) {
	one := BookingSeats{}

	query := `SELECT id,booking_id,seats_id,price,sub_total,duration_from,duration_to FROM "booking_seats" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(
		&one.ID, &one.BookingID, &one.SeatsID, &one.Price, &one.SubTotal, &one.DurationFrom, &one.DurationTo,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (b *BookingSeats) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "booking_seats" SET booking_id = $1,seats_id = $2,price = $3,sub_total = $4,duration_from = $5,duration_to = $6 WHERE id = $7 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.BookingID, b.SeatsID, b.Price, b.SubTotal, b.DurationFrom, b.DurationTo, b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (b *BookingSeats) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "booking_seats" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
