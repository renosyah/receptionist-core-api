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
	Seats struct {
		ID          uuid.UUID       `json:"id"`
		StoreID     uuid.UUID       `json:"store_id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Position    int             `json:"position"`
		Price       decimal.Decimal `json:"price"`
	}

	SeatsResponse struct {
		ID          uuid.UUID       `json:"id"`
		StoreID     uuid.UUID       `json:"store_id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Position    int             `json:"position"`
		Price       decimal.Decimal `json:"price"`
	}
)

func (s *Seats) Response() SeatsResponse {
	return SeatsResponse{
		ID:          s.ID,
		StoreID:     s.StoreID,
		Name:        s.Name,
		Description: s.Description,
		Position:    s.Position,
		Price:       s.Price,
	}
}

func (s *Seats) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "seats" (store_id,name,description,position,price) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.StoreID, s.Name, s.Description, s.Position, s.Price).Scan(&s.ID)
	if err != nil {
		return s.ID, err
	}

	return s.ID, nil
}

func (s *Seats) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Seats, error) {
	all := []Seats{}

	query, args, err := param.Query(strings.Split("id,store_id,name,description,position,price", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,store_id,name,description,position,price FROM "seats" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Seats{}
		err = rows.Scan(
			&one.ID, &one.StoreID, &one.Name, &one.Description, &one.Position, &one.Price,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (s *Seats) One(ctx context.Context, db *sql.DB) (Seats, error) {
	one := Seats{}

	query := `SELECT id,store_id,name,description,position,price FROM "seats" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID).Scan(
		&one.ID, &one.StoreID, &one.Name, &one.Description, &one.Position, &one.Price,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (s *Seats) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "seats" SET store_id = $1,name = $2,description = $3,position = $4,price = $5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.StoreID, s.Name, s.Description, s.Position, s.Price, s.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *Seats) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "seats" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
