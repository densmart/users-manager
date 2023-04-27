package rest

import (
	"errors"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

func (h *RestRouter) createUser(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.CreateUserDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.CreateUser(*h.service, data)
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

func (h *RestRouter) retrieveUser(c *gin.Context) {
	var apiErr *dto.APIError
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	result, err := usecases.RetrieveUser(*h.service, uint64(userID))
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

func (h *RestRouter) searchUsers(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.SearchUserDTO

	if err := c.ShouldBindWith(&data, binding.Form); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.RawURL = c.Request.URL.String()

	users, err := usecases.SearchUsers(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Link", users.Pagination)
	if len(users.Items) <= 0 {
		SuccessResponse(c, make([]string, 0))
		return
	}
	SuccessResponse(c, users.Items)
}

func (h *RestRouter) updateUser(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.UpdateUserDTO

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	if err = c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.ID = uint64(userID)

	user, err := usecases.UpdateUser(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, user)
}

func (h *RestRouter) deleteUser(c *gin.Context) {
	var apiErr *dto.APIError
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "user ID not specified")
		return
	}

	if err = usecases.DeleteUser(*h.service, uint64(userID)); err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	EmptyResponse(c)
}
