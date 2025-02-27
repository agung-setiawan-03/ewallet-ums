package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userID int) (*external.Wallet, error)
	SendNotification(ctx context.Context, recipient string, templateName string, placeHolder map[string]string) error
}
