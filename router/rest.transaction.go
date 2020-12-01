package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Transaction

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return transactionModule.Add(ctx, param)
}

func HandlerAllTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return transactionModule.All(ctx, param)
}

func HandlerOneTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return transactionModule.One(ctx, model.Transaction{ID: id})
}

func HandlerUpdateTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Transaction

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return transactionModule.Update(ctx, param)
}

func HandlerDeleteTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Transaction/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return transactionModule.Delete(ctx, model.Transaction{ID: id})
}
