package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddBooking(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Booking

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingModule.Add(ctx, param)
}

func HandlerAllBooking(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingModule.All(ctx, param)
}

func HandlerOneBooking(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingModule.One(ctx, model.Booking{ID: id})
}

func HandlerUpdateBooking(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Booking

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return bookingModule.Update(ctx, param)
}

func HandlerDeleteBooking(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Booking/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingModule.Delete(ctx, model.Booking{ID: id})
}
