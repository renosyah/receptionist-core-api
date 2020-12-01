package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddBookingDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.BookingDetail

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingDetailModule.Add(ctx, param)
}

func HandlerAllBookingDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingDetailModule.All(ctx, param)
}

func HandlerOneBookingDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingDetailModule.One(ctx, model.BookingDetail{ID: id})
}

func HandlerUpdateBookingDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.BookingDetail

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return bookingDetailModule.Update(ctx, param)
}

func HandlerDeleteBookingDetail(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingDetail/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingDetailModule.Delete(ctx, model.BookingDetail{ID: id})
}
