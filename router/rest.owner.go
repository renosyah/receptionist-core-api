package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Owner

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return ownerModule.Add(ctx, param)
}

func HandlerAllOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return ownerModule.All(ctx, param)
}

func HandlerOneOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return ownerModule.One(ctx, model.Owner{ID: id})
}

func HandlerLoginOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Owner

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "owner/login/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return ownerModule.Login(ctx, param)
}

func HandlerRegisterOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Owner

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "owner/register/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return ownerModule.Register(ctx, param)
}

func HandlerUpdateOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Owner

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return ownerModule.Update(ctx, param)
}

func HandlerDeleteOwner(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Owner/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return ownerModule.Delete(ctx, model.Owner{ID: id})
}
