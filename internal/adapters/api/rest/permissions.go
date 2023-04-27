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

func (h *RestRouter) createPermissions(c *gin.Context) {
	var apiErr *dto.APIError
	var data []dto.CreatePermissionDTO

	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "role ID not specified")
		return
	}

	if err = c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	for idx, _ := range data {
		data[idx].RoleID = uint64(roleID)
	}

	result, err := usecases.CreatePermissions(*h.service, data)
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

func (h *RestRouter) searchPermissions(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.SearchPermissionDTO
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "role ID not specified")
		return
	}
	if err = c.ShouldBindWith(&data, binding.Form); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.RawURL = c.Request.URL.String()
	data.RoleID = uint64(roleID)

	perms, err := usecases.SearchPermissions(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Link", perms.Pagination)
	if len(perms.Items) <= 0 {
		SuccessResponse(c, make([]string, 0))
		return
	}
	SuccessResponse(c, perms.Items)
}

func (h *RestRouter) updatePermissions(c *gin.Context) {
	var apiErr *dto.APIError
	var data []dto.UpdatePermissionDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	role, err := usecases.UpdatePermissions(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, role)
}

func (h *RestRouter) deletePermissions(c *gin.Context) {
	var apiErr *dto.APIError
	var data []uint64
	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := usecases.DeletePermissions(*h.service, data); err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	EmptyResponse(c)
}
