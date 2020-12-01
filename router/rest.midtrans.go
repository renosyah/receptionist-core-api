package router

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	"github.com/renosyah/receptionist-core-api/midtrans"
	"github.com/renosyah/receptionist-core-api/model"
)

func HandlerAddMidtransTransaction(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Midtrans

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "mintrans/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return midtransModule.Create(ctx, param)
}

func HandlerMidtransNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var param midtrans.NotificationCallback

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = midtransModule.ProcessMidtransNotification(ctx, param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
