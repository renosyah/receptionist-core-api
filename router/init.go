package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/renosyah/receptionist-core-api/api"
	mid "github.com/renosyah/receptionist-core-api/midtrans"
)

var (
	dbPool *sql.DB

	customerModule      *api.CustomerModule
	ownerModule         *api.OwnerModule
	storeModule         *api.StoreModule
	seatsModule         *api.SeatsModule
	productModule       *api.ProductModule
	bookingModule       *api.BookingModule
	bookingDetailModule *api.BookingDetailModule
	bookingSeatsModule  *api.BookingSeatsModule
	transactionModule   *api.TransactionModule
	midtransModule      *api.MidtransModule
)

func Init(db *sql.DB, p *mid.PaymentGateway) {
	dbPool = db

	customerModule = api.NewCustomerModule(db)
	ownerModule = api.NewOwnerModule(db)
	storeModule = api.NewStoreModule(db)
	seatsModule = api.NewSeatsModule(db)
	productModule = api.NewProductModule(db)
	bookingModule = api.NewBookingModule(db)
	bookingDetailModule = api.NewBookingDetailModule(db)
	transactionModule = api.NewTransactionModule(db)
	bookingSeatsModule = api.NewBookingSeatsModule(db)
	midtransModule = api.NewMidtransModule(db, p)
}

// ParseBodyData parse json-formatted request body into given struct.
func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return errors.Wrap(err, "json")
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errors.Wrap(err, "validate")
	}

	if !valid {
		return errors.New("invalid data")
	}

	return nil
}
