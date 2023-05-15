package rest

import (
	"errors"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RestRouter) signIn(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.AuthRequestDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.SignIn(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) refresh(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.RefreshRequestDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.RefreshAccessToken(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) checkOtp(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.OTPRequestDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.CheckOTPToken(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, result)
}
