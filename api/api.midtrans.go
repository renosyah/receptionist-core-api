package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/shopspring/decimal"

	"github.com/renosyah/receptionist-core-api/midtrans"
	mid "github.com/renosyah/receptionist-core-api/midtrans"
	"github.com/renosyah/receptionist-core-api/model"
)

type (
	MidtransModule struct {
		db      *sql.DB
		Payment *mid.PaymentGateway
		Name    string
	}
)

func NewMidtransModule(db *sql.DB, pay *mid.PaymentGateway) *MidtransModule {
	return &MidtransModule{
		db:      db,
		Payment: pay,
		Name:    "module/Midtrans",
	}
}

func (m MidtransModule) Create(ctx context.Context, param model.Midtrans) (mid.TransactionCreated, *Error) {

	transaction := &model.Transaction{
		BookingID:      param.BookingID,
		CustomerID:     param.CustomerID,
		Total:          decimal.NewFromInt(param.TransactionDetails.GrossAmt),
		PaymentType:    0,
		PaymentStatus:  model.TRANSACTION_STATUS_PENDING,
		PaymentOrderID: "",
		PaymentID:      "",
		PaymentTime:    "",
		ApprovalCode:   "",
	}

	transactionID, err := transaction.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Transaction"

		return mid.TransactionCreated{}, NewErrorWrap(err, m.Name, "add/Transaction",
			message, status)
	}

	param.TransactionDetails.OrderID = transactionID.String()

	created, err := m.Payment.CreateTransaction(ctx, mid.TransactionCreate{
		UserID:             param.CustomerID.String(),
		TransactionDetails: param.TransactionDetails,
		Items:              param.Items,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on create Midtrans transaction"

		return mid.TransactionCreated{}, NewErrorWrap(err, m.Name, "create/Midtrans",
			message, status)
	}

	return created, nil
}

func (m MidtransModule) ProcessMidtransNotification(ctx context.Context, param mid.NotificationCallback) error {

	id, errParse := uuid.FromString(param.OrderID)
	if errParse != nil {
		return errParse
	}

	t := &model.Transaction{ID: id}
	transaction, err := t.One(ctx, m.db)
	if err != nil {
		return errors.New("error on get one Transaction")
	}

	b := &model.Booking{ID: t.BookingID}
	bill, err := b.One(ctx, m.db)
	if err != nil {
		return errors.New("error on get one Bill")
	}

	switch param.TransactionStatus {
	case midtrans.TransactionStatusCapture, midtrans.TransactionStatusSettlement:
		transaction.PaymentStatus = model.TRANSACTION_STATUS_SUCCESS
		break
	case midtrans.TransactionStatusCancel, midtrans.TransactionStatusDenied, midtrans.TransactionStatusExpired:
		transaction.PaymentStatus = model.TRANSACTION_STATUS_FAILED
		break
	case midtrans.TransactionStatusPending:
		transaction.PaymentStatus = model.TRANSACTION_STATUS_PENDING
		break
	default:
		break
	}

	// over the counter
	transaction.CstoreCode = param.PaymentCode
	transaction.CstoreName = param.Store
	if transaction.CstoreCode != "" && transaction.CstoreName != "" {
		transaction.PaymentType = model.TRANSACTION_TYPE_OTC
	}

	transaction.PaymentOrderID = param.OrderID
	transaction.PaymentID = param.TransactionID
	transaction.PaymentTime = param.TransactionTime
	transaction.ApprovalCode = param.ApprovalCode

	// for BCA va and bni va
	// get only last data
	for _, d := range param.VaNumbers {
		transaction.BankName = d.Bank
		transaction.VA = d.VaNumber
	}

	// permata bank if still empty
	if transaction.BankName == "" && transaction.VA == "" {
		transaction.BankName = "permata bank"
		transaction.VA = param.PermataVaNumber
	}

	if transaction.BankName != "" && transaction.VA != "" {
		transaction.PaymentType = model.TRANSACTION_TYPE_TRANSFER
	}

	var emptyTUUID, emptyBUUID uuid.UUID
	it, err := transaction.Update(ctx, m.db)
	if err != nil || it == emptyTUUID {
		return errors.New("error on update transaction")
	}

	tSum := &model.SumTransactionParam{
		BillID:        bill.ID,
		PaymentStatus: model.TRANSACTION_STATUS_SUCCESS,
	}
	tsum, errSum := tSum.Sum(ctx, m.db)
	if errSum != nil {
		return errSum
	}

	if tsum.Amount.Cmp(bill.Total) >= 0 {
		bill.PaymentStatus = model.TRANSACTION_STATUS_SUCCESS
	}

	ib, err := bill.Update(ctx, m.db)
	if err != nil || ib == emptyBUUID {
		return errors.New("error on update Bill")
	}

	return nil
}
