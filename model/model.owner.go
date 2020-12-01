package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	Owner struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	OwnerResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"-"`
	}
)

func (o *Owner) Response() OwnerResponse {
	return OwnerResponse{
		ID:       o.ID,
		Name:     o.Name,
		Email:    o.Email,
		Password: o.Password,
	}
}

func (o *Owner) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "owner" (name,email,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), o.Name, o.Email, o.Password).Scan(&o.ID)
	if err != nil {
		return o.ID, err
	}

	return o.ID, nil
}

func (o *Owner) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Owner, error) {
	all := []Owner{}

	query, args, err := param.Query(strings.Split("id,name,email", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,name,email,password FROM "owner" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Owner{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email, &one.Password,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (o *Owner) One(ctx context.Context, db *sql.DB) (Owner, error) {
	one := Owner{}

	query := `SELECT id,name,email,password FROM "owner" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), o.ID).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (o *Owner) OneByEmail(ctx context.Context, db *sql.DB) (Owner, error) {
	one := Owner{}

	query := `SELECT id,name,email,password FROM "owner" WHERE email = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), o.Email).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (o *Owner) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "owner" SET name = $1, email = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), o.Name, o.Email, o.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (o *Owner) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "owner" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), o.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
