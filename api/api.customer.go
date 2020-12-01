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
	CustomerModule struct {
		db   *sql.DB
		Name string
	}
)

func NewCustomerModule(db *sql.DB) *CustomerModule {
	return &CustomerModule{
		db:   db,
		Name: "module/customer",
	}
}

func (m CustomerModule) All(ctx context.Context, param model.ListQuery) ([]model.CustomerResponse, *Error) {
	var all []model.CustomerResponse

	data, err := (&model.Customer{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Customer"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Customer found"
		}
		return []model.CustomerResponse{}, NewErrorWrap(err, m.Name, "all/Customer",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m CustomerModule) Add(ctx context.Context, param model.Customer) (model.CustomerResponse, *Error) {

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash Customer password"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "add/Customer",
			message, status)
	}

	param.Password = hash

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Customer"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "add/Customer",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m CustomerModule) One(ctx context.Context, param model.Customer) (model.CustomerResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Customer"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "one/Customer",
			message, status)
	}

	return data.Response(), nil
}

func (m CustomerModule) Login(ctx context.Context, param model.Customer) (model.CustomerAuth, *Error) {
	data, err := param.OneByPhoneNumber(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "phone number or password invalid"

		return model.CustomerAuth{}, NewErrorWrap(err, m.Name, "login/Customer",
			message, status)
	}

	err = ComparePassword(ctx, data.Password, param.Password)
	if err != nil {
		status := http.StatusUnauthorized
		message := "phone number or password invalid"

		return model.CustomerAuth{}, NewErrorWrap(errors.New(message), m.Name, "login/Customer",
			message, status)
	}

	resp := model.CustomerAuth{
		SessionID:        uuid.NewV4(),
		CustomerResponse: data.Response(),
	}

	return resp, nil
}

func (m CustomerModule) Register(ctx context.Context, param model.Customer) (model.CustomerAuth, *Error) {
	var emptyUUID uuid.UUID

	cs, errCs := param.OneByPhoneNumber(ctx, m.db)
	if errCs != nil && errors.Cause(errCs) != sql.ErrNoRows {
		status := http.StatusInternalServerError
		message := "error on get one Customer"

		return model.CustomerAuth{}, NewErrorWrap(errCs, m.Name, "register/Customer",
			message, status)
	}

	if cs.ID != emptyUUID {
		status := http.StatusInternalServerError
		message := "Customer with this phone number exist"

		return model.CustomerAuth{}, NewErrorWrap(errors.New(message), m.Name, "register/Customer",
			message, status)
	}

	data, err := m.Add(ctx, param)
	if err != nil {
		return model.CustomerAuth{}, err
	}

	resp := model.CustomerAuth{
		SessionID:        uuid.NewV4(),
		CustomerResponse: data,
	}

	return resp, nil
}

func (m CustomerModule) Update(ctx context.Context, param model.Customer) (model.CustomerResponse, *Error) {
	var emptyUUID uuid.UUID

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash Customer password"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "update/Customer",
			message, status)
	}

	param.Password = hash

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Customer"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "update/Customer",
			message, status)
	}
	return param.Response(), nil
}

func (m CustomerModule) Delete(ctx context.Context, param model.Customer) (model.CustomerResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Customer"

		return model.CustomerResponse{}, NewErrorWrap(err, m.Name, "delete/Customer",
			message, status)
	}
	return param.Response(), nil
}
