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
	BookingDetail struct {
		ID        uuid.UUID       `json:"id"`
		BookingID uuid.UUID       `json:"booking_id"`
		ProductID uuid.UUID       `json:"product_id"`
		Price     decimal.Decimal `json:"price"`
		Quantity  int             `json:"quantity"`
		SubTotal  decimal.Decimal `json:"sub_total"`
	}

	BookingDetailResponse struct {
		ID        uuid.UUID       `json:"id"`
		BookingID uuid.UUID       `json:"booking_id"`
		ProductID uuid.UUID       `json:"product_id"`
		Price     decimal.Decimal `json:"price"`
		Quantity  int             `json:"quantity"`
		SubTotal  decimal.Decimal `json:"sub_total"`
	}
)

func (b *BookingDetail) Response() BookingDetailResponse {
	return BookingDetailResponse{
		ID:        b.ID,
		BookingID: b.BookingID,
		ProductID: b.ProductID,
		Price:     b.Price,
		Quantity:  b.Quantity,
		SubTotal:  b.SubTotal,
	}
}

func (b *BookingDetail) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "booking_detail" (booking_id,product_id,price,quantity,sub_total) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.BookingID, b.ProductID, b.Price, b.Quantity, b.SubTotal).Scan(&b.ID)
	if err != nil {
		return b.ID, err
	}

	return b.ID, nil
}

func (b *BookingDetail) All(ctx context.Context, db *sql.DB, param ListQuery) ([]BookingDetail, error) {
	all := []BookingDetail{}

	query, args, err := param.Query(strings.Split("id,booking_id,product_id,price,quantity,sub_total", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,booking_id,product_id,price,quantity,sub_total FROM "booking_detail" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := BookingDetail{}
		err = rows.Scan(
			&one.ID, &one.BookingID, &one.ProductID, &one.Price, &one.Quantity, &one.SubTotal,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (b *BookingDetail) One(ctx context.Context, db *sql.DB) (BookingDetail, error) {
	one := BookingDetail{}

	query := `SELECT id,booking_id,product_id,price,quantity,sub_total FROM "booking_detail" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(
		&one.ID, &one.BookingID, &one.ProductID, &one.Price, &one.Quantity, &one.SubTotal,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (b *BookingDetail) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "booking_detail" SET booking_id = $1,product_id = $2,price = $3,quantity = $4,sub_total = $5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.BookingID, b.ProductID, b.Price, b.Quantity, b.SubTotal, b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (b *BookingDetail) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "booking_detail" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
