package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddBookingSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.BookingSeats

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingSeatsModule.Add(ctx, param)
}

func HandlerAllBookingSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return bookingSeatsModule.All(ctx, param)
}

func HandlerOneBookingSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingSeatsModule.One(ctx, model.BookingSeats{ID: id})
}

func HandlerUpdateBookingSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.BookingSeats

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return bookingSeatsModule.Update(ctx, param)
}

func HandlerDeleteBookingSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "BookingSeats/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return bookingSeatsModule.Delete(ctx, model.BookingSeats{ID: id})
}
