package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Customer

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return customerModule.Add(ctx, param)
}

func HandlerAllCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return customerModule.All(ctx, param)
}

func HandlerOneCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return customerModule.One(ctx, model.Customer{ID: id})
}

func HandlerLoginCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Customer

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/login/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return customerModule.Login(ctx, param)
}

func HandlerRegisterCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Customer

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/register/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return customerModule.Register(ctx, param)
}

func HandlerUpdateCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Customer

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return customerModule.Update(ctx, param)
}

func HandlerDeleteCustomer(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "customer/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return customerModule.Delete(ctx, model.Customer{ID: id})
}
