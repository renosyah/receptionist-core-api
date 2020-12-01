package midtrans

import (
	"context"

	mid "github.com/veritrans/go-midtrans"
)

type (
	PaymentGateway struct {
		ServerKey string `json:"-"`
		ClientKey string `json:"client_key"`
		Env       string `json:"-"`
		CallBack  string `json:"callback"`
	}

	TransactionCreate struct {
		UserID             string                 `json:"user_id"`
		TransactionDetails mid.TransactionDetails `json:"transaction"`
		Items              []mid.ItemDetail       `json:"items"`
	}

	TransactionCreated struct {
		SnapToken   string `json:"snap_token"`
		RedirectURL string `json:"redirect_url"`
	}
)

func (p *PaymentGateway) CreateTransaction(ctx context.Context, create TransactionCreate) (TransactionCreated, error) {

	created := TransactionCreated{}

	midclient := mid.NewClient()
	midclient.ServerKey = p.ServerKey
	midclient.ClientKey = p.ClientKey

	switch p.Env {
	case "sanbox":
		midclient.APIEnvType = mid.Sandbox
		break
	case "production":
		midclient.APIEnvType = mid.Production
		break
	default:
		break
	}

	var snapGateway mid.SnapGateway
	snapGateway = mid.SnapGateway{
		Client: midclient,
	}

	snapReq := &mid.SnapReq{
		TransactionDetails: create.TransactionDetails,
		Items:              &create.Items,
		UserId:             create.UserID,
		Callbacks: &mid.Callbacks{
			Finish: p.CallBack,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return created, err
	}
	created.SnapToken = snapTokenResp.Token
	created.RedirectURL = snapTokenResp.RedirectURL

	return created, nil
}
