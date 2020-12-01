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
	Product struct {
		ID          uuid.UUID       `json:"id"`
		StoreID     uuid.UUID       `json:"store_id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		ImageURL    string          `json:"image_url"`
		Price       decimal.Decimal `json:"price"`
	}

	ProductResponse struct {
		ID          uuid.UUID       `json:"id"`
		StoreID     uuid.UUID       `json:"store_id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		ImageURL    string          `json:"image_url"`
		Price       decimal.Decimal `json:"price"`
	}
)

func (p *Product) Response() ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		StoreID:     p.StoreID,
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		Price:       p.Price,
	}
}

func (p *Product) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "product" (store_id,name,description,image_url,price) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.StoreID, p.Name, p.Description, p.ImageURL, p.Price).Scan(&p.ID)
	if err != nil {
		return p.ID, err
	}

	return p.ID, nil
}

func (p *Product) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Product, error) {
	all := []Product{}

	query, args, err := param.Query(strings.Split("id,store_id,name,description,image_url,price", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,store_id,name,description,image_url,price FROM "product" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Product{}
		err = rows.Scan(
			&one.ID, &one.StoreID, &one.Name, &one.Description, &one.ImageURL, &one.Price,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (p *Product) One(ctx context.Context, db *sql.DB) (Product, error) {
	one := Product{}

	query := `SELECT id,store_id,name,description,image_url,price FROM "product" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.ID).Scan(
		&one.ID, &one.StoreID, &one.Name, &one.Description, &one.ImageURL, &one.Price,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (p *Product) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "product" SET store_id = $1,name = $2,description = $3,image_url = $4,price = $5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.StoreID, p.Name, p.Description, p.ImageURL, p.Price, p.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (p *Product) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "product" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), p.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
