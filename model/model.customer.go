package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	Customer struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		PhoneNumber string    `json:"phone_number"`
		Password    string    `json:"password"`
	}

	CustomerResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		PhoneNumber string    `json:"phone_number"`
		Password    string    `json:"-"`
	}
)

func (c *Customer) Response() CustomerResponse {
	return CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Password:    c.Password,
	}
}

func (c *Customer) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "customer" (name,phone_number,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.Name, c.PhoneNumber, c.Password).Scan(&c.ID)
	if err != nil {
		return c.ID, err
	}

	return c.ID, nil
}

func (c *Customer) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Customer, error) {
	all := []Customer{}

	query, args, err := param.Query(strings.Split("id,name,phone_number", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,name,phone_number,password FROM "customer" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Customer{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (c *Customer) One(ctx context.Context, db *sql.DB) (Customer, error) {
	one := Customer{}

	query := `SELECT id,name,phone_number,password FROM "customer" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (c *Customer) OneByPhoneNumber(ctx context.Context, db *sql.DB) (Customer, error) {
	one := Customer{}

	query := `SELECT id,name,phone_number,password FROM "customer" WHERE phone_number = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.PhoneNumber).Scan(
		&one.ID, &one.Name, &one.PhoneNumber, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (c *Customer) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "customer" SET name = $1, phone_number = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.Name, c.PhoneNumber, c.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (c *Customer) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "customer" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
