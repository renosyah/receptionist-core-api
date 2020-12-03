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
	BookingSeatsModule struct {
		db   *sql.DB
		Name string
	}
)

func NewBookingSeatsModule(db *sql.DB) *BookingSeatsModule {
	return &BookingSeatsModule{
		db:   db,
		Name: "module/BookingSeats",
	}
}

func (m BookingSeatsModule) All(ctx context.Context, param model.ListQuery) ([]model.BookingSeatsResponse, *Error) {
	var all []model.BookingSeatsResponse

	data, err := (&model.BookingSeats{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all BookingSeats"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no BookingSeats found"
		}
		return []model.BookingSeatsResponse{}, NewErrorWrap(err, m.Name, "all/BookingSeats",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m BookingSeatsModule) Add(ctx context.Context, param model.BookingSeats) (model.BookingSeatsResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add BookingSeats"

		return model.BookingSeatsResponse{}, NewErrorWrap(err, m.Name, "add/BookingSeats",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m BookingSeatsModule) One(ctx context.Context, param model.BookingSeats) (model.BookingSeatsResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one BookingSeats"

		return model.BookingSeatsResponse{}, NewErrorWrap(err, m.Name, "one/BookingSeats",
			message, status)
	}

	return data.Response(), nil
}

func (m BookingSeatsModule) Update(ctx context.Context, param model.BookingSeats) (model.BookingSeatsResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update BookingSeats"

		return model.BookingSeatsResponse{}, NewErrorWrap(err, m.Name, "update/BookingSeats",
			message, status)
	}
	return param.Response(), nil
}

func (m BookingSeatsModule) Delete(ctx context.Context, param model.BookingSeats) (model.BookingSeatsResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete BookingSeats"

		return model.BookingSeatsResponse{}, NewErrorWrap(err, m.Name, "delete/BookingSeats",
			message, status)
	}
	return param.Response(), nil
}
