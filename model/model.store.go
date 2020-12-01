package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	Store struct {
		ID          uuid.UUID `json:"id"`
		OwnerID     uuid.UUID `json:"owner_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ImageURL    string    `json:"image_url"`
		Latitude    float64   `json:"latitude"`
		Longitude   float64   `json:"longitude"`
	}

	StoreResponse struct {
		ID          uuid.UUID `json:"id"`
		OwnerID     uuid.UUID `json:"owner_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ImageURL    string    `json:"image_url"`
		Latitude    float64   `json:"latitude"`
		Longitude   float64   `json:"longitude"`
	}
)

func (s *Store) Response() StoreResponse {
	return StoreResponse{
		ID:          s.ID,
		OwnerID:     s.OwnerID,
		Name:        s.Name,
		Description: s.Description,
		ImageURL:    s.ImageURL,
		Latitude:    s.Latitude,
		Longitude:   s.Longitude,
	}
}

func (s *Store) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "store" (owner_id,name,description,image_url,latitude,longitude) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.OwnerID, s.Name, s.Description, s.ImageURL, s.Latitude, s.Longitude).Scan(&s.ID)
	if err != nil {
		return s.ID, err
	}

	return s.ID, nil
}

func (s *Store) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Store, error) {
	all := []Store{}

	query, args, err := param.Query(strings.Split("id,owner_id,name,description,image_url,latitude,longitude", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,owner_id,name,description,image_url,latitude,longitude FROM "store" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Store{}
		err = rows.Scan(
			&one.ID, &one.OwnerID, &one.Name, &one.Description, &one.ImageURL, &one.Latitude, &one.Longitude,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (s *Store) One(ctx context.Context, db *sql.DB) (Store, error) {
	one := Store{}

	query := `SELECT id,owner_id,name,description,image_url,latitude,longitude FROM "store" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID).Scan(
		&one.ID, &one.OwnerID, &one.Name, &one.Description, &one.ImageURL, &one.Latitude, &one.Longitude,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (s *Store) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "store" SET owner_id = $1,name = $2,description = $3,image_url = $4,latitude = $5,longitude = $6 WHERE id = $7 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.OwnerID, s.Name, s.Description, s.ImageURL, s.Latitude, s.Longitude, s.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *Store) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "store" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
