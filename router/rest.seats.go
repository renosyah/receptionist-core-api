package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Seats

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return seatsModule.Add(ctx, param)
}

func HandlerAllSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return seatsModule.All(ctx, param)
}

func HandlerOneSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return seatsModule.One(ctx, model.Seats{ID: id})
}

func HandlerUpdateSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Seats

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return seatsModule.Update(ctx, param)
}

func HandlerDeleteSeats(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Seats/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return seatsModule.Delete(ctx, model.Seats{ID: id})
}
