package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	decimal "github.com/shopspring/decimal"
)

const (
	TRANSACTION_STATUS_SUCCESS = 0
	TRANSACTION_STATUS_PENDING = 1
	TRANSACTION_STATUS_FAILED  = 2

	TRANSACTION_TYPE_TRANSFER = 1
	TRANSACTION_TYPE_OTC      = 2
)

type (
	SumTransactionParam struct {
		BillID        uuid.UUID `json:"bill_id"`
		PaymentStatus int       `json:"payment_status"`
	}

	SumTransaction struct {
		Amount decimal.Decimal `json:"total"`
	}

	Transaction struct {
		ID             uuid.UUID       `json:"id"`
		BookingID      uuid.UUID       `json:"booking_id"`
		CustomerID     uuid.UUID       `json:"customer_id"`
		Total          decimal.Decimal `json:"total"`
		PaymentType    int             `json:"payment_type"`
		PaymentStatus  int             `json:"payment_status"`
		PaymentOrderID string          `json:"payment_order_id"`
		PaymentID      string          `json:"payment_id"`
		PaymentTime    string          `json:"payment_time"`
		ApprovalCode   string          `json:"approval_code"`
		BankName       string          `json:"bank_name"`
		VA             string          `json:"va"`
		CstoreCode     string          `json:"cstore_code"`
		CstoreName     string          `json:"cstore_name"`
	}

	TransactionResponse struct {
		ID             uuid.UUID       `json:"id"`
		BookingID      uuid.UUID       `json:"booking_id"`
		CustomerID     uuid.UUID       `json:"customer_id"`
		Total          decimal.Decimal `json:"total"`
		PaymentType    int             `json:"payment_type"`
		PaymentStatus  int             `json:"payment_status"`
		PaymentOrderID string          `json:"payment_order_id"`
		PaymentID      string          `json:"payment_id"`
		PaymentTime    string          `json:"payment_time"`
		ApprovalCode   string          `json:"approval_code"`
		BankName       string          `json:"bank_name"`
		VA             string          `json:"va"`
		CstoreCode     string          `json:"cstore_code"`
		CstoreName     string          `json:"cstore_name"`
	}
)

func (t *Transaction) Response() TransactionResponse {
	return TransactionResponse{
		ID:             t.ID,
		BookingID:      t.BookingID,
		CustomerID:     t.CustomerID,
		Total:          t.Total,
		PaymentType:    t.PaymentType,
		PaymentStatus:  t.PaymentStatus,
		PaymentOrderID: t.PaymentOrderID,
		PaymentID:      t.PaymentID,
		PaymentTime:    t.PaymentTime,
		ApprovalCode:   t.ApprovalCode,
		BankName:       t.BankName,
		VA:             t.VA,
		CstoreCode:     t.CstoreCode,
		CstoreName:     t.CstoreName,
	}
}

func (t *Transaction) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "transaction" (booking_id,customer_id,total,payment_type,payment_status,payment_order_id,payment_id,payment_time,approval_code,bank_name,va,cstore_code,cstore_name) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query),
		t.BookingID,
		t.CustomerID,
		t.Total,
		t.PaymentType,
		t.PaymentStatus,
		t.PaymentOrderID,
		t.PaymentID,
		t.PaymentTime,
		t.ApprovalCode,
		t.BankName,
		t.VA,
		t.CstoreCode,
		t.CstoreName,
	).Scan(&t.ID)
	if err != nil {
		return t.ID, err
	}

	return t.ID, nil
}

func (t *Transaction) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Transaction, error) {
	all := []Transaction{}

	query, args, err := param.Query(strings.Split("id,booking_id,customer_id,total,payment_type,payment_status,payment_order_id,payment_id,payment_time,approval_code,bank_name,va,cstore_code,cstore_name", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,booking_id,customer_id,total,payment_type,payment_status,payment_order_id,payment_id,payment_time,approval_code,bank_name,va,cstore_code,cstore_name FROM "transaction" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Transaction{}
		err = rows.Scan(
			&one.ID,
			&one.BookingID,
			&one.CustomerID,
			&one.Total,
			&one.PaymentType,
			&one.PaymentStatus,
			&one.PaymentOrderID,
			&one.PaymentID,
			&one.PaymentTime,
			&one.ApprovalCode,
			&one.BankName,
			&one.VA,
			&one.CstoreCode,
			&one.CstoreName,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (t *Transaction) One(ctx context.Context, db *sql.DB) (Transaction, error) {
	one := Transaction{}

	query := `SELECT id,booking_id,customer_id,total,payment_type,payment_status,payment_order_id,payment_id,payment_time,approval_code,bank_name,va,cstore_code,cstore_name FROM "transaction" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.ID).Scan(
		&one.ID,
		&one.BookingID,
		&one.CustomerID,
		&one.Total,
		&one.PaymentType,
		&one.PaymentStatus,
		&one.PaymentOrderID,
		&one.PaymentID,
		&one.PaymentTime,
		&one.ApprovalCode,
		&one.BankName,
		&one.VA,
		&one.CstoreCode,
		&one.CstoreName,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *SumTransactionParam) Sum(ctx context.Context, db *sql.DB) (SumTransaction, error) {
	one := SumTransaction{}

	query := `SELECT COALESCE(SUM(amount),0) FROM "transaction" WHERE bill_id = $1 AND payment_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.BillID, u.PaymentStatus).Scan(
		&one.Amount,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (t *Transaction) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "transaction" SET booking_id = $1,customer_id = $2,total = $3,payment_type = $4,payment_status = $5,payment_order_id = $6,payment_id = $7,payment_time = $8,approval_code = $9,bank_name = $10,va = $11,cstore_code = $10,cstore_name = $11 WHERE id = $12 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query),
		t.BookingID,
		t.CustomerID,
		t.Total,
		t.PaymentType,
		t.PaymentStatus,
		t.PaymentOrderID,
		t.PaymentID,
		t.PaymentTime,
		t.ApprovalCode,
		t.BankName,
		t.VA,
		t.CstoreCode,
		t.CstoreName,
		t.ID,
	).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (t *Transaction) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "transaction" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
