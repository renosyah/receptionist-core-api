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
	OwnerModule struct {
		db   *sql.DB
		Name string
	}
)

func NewOwnerModule(db *sql.DB) *OwnerModule {
	return &OwnerModule{
		db:   db,
		Name: "module/Owner",
	}
}

func (m OwnerModule) All(ctx context.Context, param model.ListQuery) ([]model.OwnerResponse, *Error) {
	var all []model.OwnerResponse

	data, err := (&model.Owner{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Owner"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Owner found"
		}
		return []model.OwnerResponse{}, NewErrorWrap(err, m.Name, "all/Owner",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m OwnerModule) Add(ctx context.Context, param model.Owner) (model.OwnerResponse, *Error) {

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash Owner password"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "add/Owner",
			message, status)
	}

	param.Password = hash

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Owner"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "add/Owner",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m OwnerModule) One(ctx context.Context, param model.Owner) (model.OwnerResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Owner"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "one/Owner",
			message, status)
	}

	return data.Response(), nil
}

func (m OwnerModule) Login(ctx context.Context, param model.Owner) (model.OwnerAuth, *Error) {
	data, err := param.OneByEmail(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "email or password invalid"

		return model.OwnerAuth{}, NewErrorWrap(err, m.Name, "login/Owner",
			message, status)
	}

	err = ComparePassword(ctx, data.Password, param.Password)
	if err != nil {
		status := http.StatusUnauthorized
		message := "email or password invalid"

		return model.OwnerAuth{}, NewErrorWrap(errors.New(message), m.Name, "login/Owner",
			message, status)
	}

	resp := model.OwnerAuth{
		SessionID:     uuid.NewV4(),
		OwnerResponse: data.Response(),
	}

	return resp, nil
}

func (m OwnerModule) Register(ctx context.Context, param model.Owner) (model.OwnerAuth, *Error) {
	var emptyUUID uuid.UUID

	owr, errOwr := param.OneByEmail(ctx, m.db)
	if errOwr != nil && errors.Cause(errOwr) != sql.ErrNoRows {
		status := http.StatusInternalServerError
		message := "error on get one Owner"

		return model.OwnerAuth{}, NewErrorWrap(errOwr, m.Name, "register/Owner",
			message, status)
	}

	if owr.ID != emptyUUID {
		status := http.StatusInternalServerError
		message := "Owner with this email exist"

		return model.OwnerAuth{}, NewErrorWrap(errors.New(message), m.Name, "register/Owner",
			message, status)
	}

	data, err := m.Add(ctx, param)
	if err != nil {
		return model.OwnerAuth{}, err
	}

	resp := model.OwnerAuth{
		SessionID:     uuid.NewV4(),
		OwnerResponse: data,
	}

	return resp, nil
}

func (m OwnerModule) Update(ctx context.Context, param model.Owner) (model.OwnerResponse, *Error) {
	var emptyUUID uuid.UUID

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash Owner password"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "update/Owner",
			message, status)
	}

	param.Password = hash

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Owner"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "update/Owner",
			message, status)
	}
	return param.Response(), nil
}

func (m OwnerModule) Delete(ctx context.Context, param model.Owner) (model.OwnerResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Owner"

		return model.OwnerResponse{}, NewErrorWrap(err, m.Name, "delete/Owner",
			message, status)
	}
	return param.Response(), nil
}
