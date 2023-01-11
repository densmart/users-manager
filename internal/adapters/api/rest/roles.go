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

func (h *RestRouter) createRole(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.CreateRoleDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.CreateRole(*h.service, data)
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

func (h *RestRouter) retrieveRole(c *gin.Context) {
	var apiErr *dto.APIError
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	result, err := usecases.RetrieveRole(*h.service, uint64(roleID))
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) searchRoles(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.SearchRoleDTO

	if err := c.ShouldBindWith(&data, binding.Form); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.RawURL = c.Request.URL.String()

	roles, err := usecases.SearchRoles(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Link", roles.Pagination)
	SuccessResponse(c, roles)
}

func (h *RestRouter) updateRole(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.UpdateRoleDTO

	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	if err = c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.ID = uint64(roleID)

	role, err := usecases.UpdateRole(*h.service, data)
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

func (h *RestRouter) deleteRole(c *gin.Context) {
	var apiErr *dto.APIError
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "role ID not specified")
		return
	}

	if err = usecases.DeleteRole(*h.service, uint64(roleID)); err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	EmptyResponse(c)
}
