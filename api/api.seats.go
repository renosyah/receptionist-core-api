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
	SeatsModule struct {
		db   *sql.DB
		Name string
	}
)

func NewSeatsModule(db *sql.DB) *SeatsModule {
	return &SeatsModule{
		db:   db,
		Name: "module/Seats",
	}
}

func (m SeatsModule) All(ctx context.Context, param model.ListQuery) ([]model.SeatsResponse, *Error) {
	var all []model.SeatsResponse

	data, err := (&model.Seats{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Seats"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Seats found"
		}
		return []model.SeatsResponse{}, NewErrorWrap(err, m.Name, "all/Seats",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m SeatsModule) Add(ctx context.Context, param model.Seats) (model.SeatsResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Seats"

		return model.SeatsResponse{}, NewErrorWrap(err, m.Name, "add/Seats",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m SeatsModule) One(ctx context.Context, param model.Seats) (model.SeatsResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Seats"

		return model.SeatsResponse{}, NewErrorWrap(err, m.Name, "one/Seats",
			message, status)
	}

	return data.Response(), nil
}

func (m SeatsModule) Update(ctx context.Context, param model.Seats) (model.SeatsResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Seats"

		return model.SeatsResponse{}, NewErrorWrap(err, m.Name, "update/Seats",
			message, status)
	}
	return param.Response(), nil
}

func (m SeatsModule) Delete(ctx context.Context, param model.Seats) (model.SeatsResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Seats"

		return model.SeatsResponse{}, NewErrorWrap(err, m.Name, "delete/Seats",
			message, status)
	}
	return param.Response(), nil
}
