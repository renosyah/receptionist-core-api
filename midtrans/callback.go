package midtrans

const TransactionStatusCapture = "capture"
const TransactionStatusSettlement = "settlement"
const TransactionStatusCancel = "cancel"
const TransactionStatusDenied = "deny"
const TransactionStatusExpired = "expire"
const TransactionStatusPending = "pending"

const FraudStatusChallenge = "challenge"
const FraudStatusAccept = "accept"

const PaymentTypeCstore = "cstore"
const PaymentTypeBankTransferVA = "bank_transfer"

type NotificationCallback struct {

	// permata VA
	PermataVaNumber string `json:"permata_va_number"`

	// bni & bca VA
	VaNumbers []struct {
		Bank     string `json:"bank"`
		VaNumber string `json:"va_number"`
	} `json:"va_numbers"`
	PaymentAmounts []struct {
		PaidAt string `json:"paid_at"`
		Amount string `json:"amount"`
	} `json:"payment_amounts"`

	// over the counter
	PaymentCode string `json:"payment_code"`
	Store       string `json:"store"`

	// common
	TransactionTime        string `json:"transaction_time"`
	TransactionStatus      string `json:"transaction_status"`
	TransactionID          string `json:"transaction_id"`
	StatusMessage          string `json:"status_message"`
	StatusCode             string `json:"status_code"`
	SignatureKey           string `json:"signature_key"`
	OrderID                string `json:"order_id"`
	MerchantID             string `json:"merchant_id"`
	MaskedCard             string `json:"masked_card"`
	GrossAmount            string `json:"gross_amount"`
	FraudStatus            string `json:"fraud_status"`
	Eci                    string `json:"eci"`
	Currency               string `json:"currency"`
	ChannelResponseMessage string `json:"channel_response_message"`
	ChannelResponseCode    string `json:"channel_response_code"`
	CardType               string `json:"card_type"`
	Bank                   string `json:"bank"`
	ApprovalCode           string `json:"approval_code"`
	PaymentType            string `json:"payment_type"`
}
