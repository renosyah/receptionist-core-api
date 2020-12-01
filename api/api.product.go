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
	ProductModule struct {
		db   *sql.DB
		Name string
	}
)

func NewProductModule(db *sql.DB) *ProductModule {
	return &ProductModule{
		db:   db,
		Name: "module/Product",
	}
}

func (m ProductModule) All(ctx context.Context, param model.ListQuery) ([]model.ProductResponse, *Error) {
	var all []model.ProductResponse

	data, err := (&model.Product{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Product"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Product found"
		}
		return []model.ProductResponse{}, NewErrorWrap(err, m.Name, "all/Product",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m ProductModule) Add(ctx context.Context, param model.Product) (model.ProductResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Product"

		return model.ProductResponse{}, NewErrorWrap(err, m.Name, "add/Product",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m ProductModule) One(ctx context.Context, param model.Product) (model.ProductResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Product"

		return model.ProductResponse{}, NewErrorWrap(err, m.Name, "one/Product",
			message, status)
	}

	return data.Response(), nil
}

func (m ProductModule) Update(ctx context.Context, param model.Product) (model.ProductResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Product"

		return model.ProductResponse{}, NewErrorWrap(err, m.Name, "update/Product",
			message, status)
	}
	return param.Response(), nil
}

func (m ProductModule) Delete(ctx context.Context, param model.Product) (model.ProductResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Product"

		return model.ProductResponse{}, NewErrorWrap(err, m.Name, "delete/Product",
			message, status)
	}
	return param.Response(), nil
}
