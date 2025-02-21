package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}

type IRegisterHandler interface {
	Register(*gin.Context)
}
