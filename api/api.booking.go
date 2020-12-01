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
	BookingModule struct {
		db   *sql.DB
		Name string
	}
)

func NewBookingModule(db *sql.DB) *BookingModule {
	return &BookingModule{
		db:   db,
		Name: "module/Booking",
	}
}

func (m BookingModule) All(ctx context.Context, param model.ListQuery) ([]model.BookingResponse, *Error) {
	var all []model.BookingResponse

	data, err := (&model.Booking{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Booking"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Booking found"
		}
		return []model.BookingResponse{}, NewErrorWrap(err, m.Name, "all/Booking",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m BookingModule) Add(ctx context.Context, param model.Booking) (model.BookingResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Booking"

		return model.BookingResponse{}, NewErrorWrap(err, m.Name, "add/Booking",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m BookingModule) One(ctx context.Context, param model.Booking) (model.BookingResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Booking"

		return model.BookingResponse{}, NewErrorWrap(err, m.Name, "one/Booking",
			message, status)
	}

	return data.Response(), nil
}

func (m BookingModule) Update(ctx context.Context, param model.Booking) (model.BookingResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Booking"

		return model.BookingResponse{}, NewErrorWrap(err, m.Name, "update/Booking",
			message, status)
	}
	return param.Response(), nil
}

func (m BookingModule) Delete(ctx context.Context, param model.Booking) (model.BookingResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Booking"

		return model.BookingResponse{}, NewErrorWrap(err, m.Name, "delete/Booking",
			message, status)
	}
	return param.Response(), nil
}
