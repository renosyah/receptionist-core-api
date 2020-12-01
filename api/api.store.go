package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

type (
	StoreModule struct {
		db   *sql.DB
		Name string
	}
)

func NewStoreModule(db *sql.DB) *StoreModule {
	return &StoreModule{
		db:   db,
		Name: "module/Store",
	}
}

func (m StoreModule) All(ctx context.Context, param model.ListQuery) ([]model.StoreResponse, *Error) {
	var all []model.StoreResponse

	data, err := (&model.Store{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Store"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Store found"
		}
		return []model.StoreResponse{}, NewErrorWrap(err, m.Name, "all/Store",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m StoreModule) Add(ctx context.Context, param model.Store) (model.StoreResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Store"

		return model.StoreResponse{}, NewErrorWrap(err, m.Name, "add/Store",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m StoreModule) One(ctx context.Context, param model.Store) (model.StoreResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Store"

		return model.StoreResponse{}, NewErrorWrap(err, m.Name, "one/Store",
			message, status)
	}

	return data.Response(), nil
}

func (m StoreModule) Update(ctx context.Context, param model.Store) (model.StoreResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Store"

		return model.StoreResponse{}, NewErrorWrap(err, m.Name, "update/Store",
			message, status)
	}
	return param.Response(), nil
}

func (m StoreModule) Delete(ctx context.Context, param model.Store) (model.StoreResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Store"

		return model.StoreResponse{}, NewErrorWrap(err, m.Name, "delete/Store",
			message, status)
	}
	return param.Response(), nil
}
