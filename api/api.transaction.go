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
	TransactionModule struct {
		db   *sql.DB
		Name string
	}
)

func NewTransactionModule(db *sql.DB) *TransactionModule {
	return &TransactionModule{
		db:   db,
		Name: "module/Transaction",
	}
}

func (m TransactionModule) All(ctx context.Context, param model.ListQuery) ([]model.TransactionResponse, *Error) {
	var all []model.TransactionResponse

	data, err := (&model.Transaction{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Transaction"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Transaction found"
		}
		return []model.TransactionResponse{}, NewErrorWrap(err, m.Name, "all/Transaction",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m TransactionModule) Add(ctx context.Context, param model.Transaction) (model.TransactionResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Transaction"

		return model.TransactionResponse{}, NewErrorWrap(err, m.Name, "add/Transaction",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m TransactionModule) Sum(ctx context.Context, param model.SumTransactionParam) (model.SumTransaction, *Error) {
	data, err := param.Sum(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get sum Transaction"

		return model.SumTransaction{}, NewErrorWrap(err, m.Name, "sum/Transaction",
			message, status)
	}

	return data, nil
}

func (m TransactionModule) One(ctx context.Context, param model.Transaction) (model.TransactionResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Transaction"

		return model.TransactionResponse{}, NewErrorWrap(err, m.Name, "one/Transaction",
			message, status)
	}

	return data.Response(), nil
}

func (m TransactionModule) Update(ctx context.Context, param model.Transaction) (model.TransactionResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Transaction"

		return model.TransactionResponse{}, NewErrorWrap(err, m.Name, "update/Transaction",
			message, status)
	}
	return param.Response(), nil
}

func (m TransactionModule) Delete(ctx context.Context, param model.Transaction) (model.TransactionResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Transaction"

		return model.TransactionResponse{}, NewErrorWrap(err, m.Name, "delete/Transaction",
			message, status)
	}
	return param.Response(), nil
}
