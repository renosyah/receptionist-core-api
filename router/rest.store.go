package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddStore(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Store

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return storeModule.Add(ctx, param)
}

func HandlerAllStore(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return storeModule.All(ctx, param)
}

func HandlerOneStore(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return storeModule.One(ctx, model.Store{ID: id})
}

func HandlerUpdateStore(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Store

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return storeModule.Update(ctx, param)
}

func HandlerDeleteStore(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Store/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return storeModule.Delete(ctx, model.Store{ID: id})
}
