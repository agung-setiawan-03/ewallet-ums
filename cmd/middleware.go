package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("authorization")
	if auth == "" {
		log.Println("Authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("Failed to get user session on DB: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}