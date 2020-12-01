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
	BookingDetailModule struct {
		db   *sql.DB
		Name string
	}
)

func NewBookingDetailModule(db *sql.DB) *BookingDetailModule {
	return &BookingDetailModule{
		db:   db,
		Name: "module/BookingDetail",
	}
}

func (m BookingDetailModule) All(ctx context.Context, param model.ListQuery) ([]model.BookingDetailResponse, *Error) {
	var all []model.BookingDetailResponse

	data, err := (&model.BookingDetail{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all BookingDetail"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no BookingDetail found"
		}
		return []model.BookingDetailResponse{}, NewErrorWrap(err, m.Name, "all/BookingDetail",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m BookingDetailModule) Add(ctx context.Context, param model.BookingDetail) (model.BookingDetailResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add BookingDetail"

		return model.BookingDetailResponse{}, NewErrorWrap(err, m.Name, "add/BookingDetail",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m BookingDetailModule) One(ctx context.Context, param model.BookingDetail) (model.BookingDetailResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one BookingDetail"

		return model.BookingDetailResponse{}, NewErrorWrap(err, m.Name, "one/BookingDetail",
			message, status)
	}

	return data.Response(), nil
}

func (m BookingDetailModule) Update(ctx context.Context, param model.BookingDetail) (model.BookingDetailResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update BookingDetail"

		return model.BookingDetailResponse{}, NewErrorWrap(err, m.Name, "update/BookingDetail",
			message, status)
	}
	return param.Response(), nil
}

func (m BookingDetailModule) Delete(ctx context.Context, param model.BookingDetail) (model.BookingDetailResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete BookingDetail"

		return model.BookingDetailResponse{}, NewErrorWrap(err, m.Name, "delete/BookingDetail",
			message, status)
	}
	return param.Response(), nil
}
