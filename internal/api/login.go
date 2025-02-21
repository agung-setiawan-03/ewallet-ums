package api

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.LoginRequest
		resp models.LoginResponse
	)

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error("Failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest,constants.ErrFailedBadRequest, nil)
		return 
	}

	resp, err := api.LoginService.Login(c.Request.Context(), req)

	if err != nil {
		log.Error("Failed on login service: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
	
}
