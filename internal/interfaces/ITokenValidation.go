package interfaces

import (
	"context"
	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
