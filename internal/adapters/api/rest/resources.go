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

func (h *RestRouter) createResource(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.CreateResourceDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.CreateResource(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) retrieveResource(c *gin.Context) {
	var apiErr *dto.APIError
	resourceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	result, err := usecases.RetrieveResource(*h.service, uint64(resourceID))
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) searchResources(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.SearchResourceDTO

	if err := c.ShouldBindWith(&data, binding.Form); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.RawURL = c.Request.URL.String()

	resources, err := usecases.SearchResources(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Link", resources.Pagination)
	SuccessResponse(c, resources)
}

func (h *RestRouter) updateResource(c *gin.Context) {
	var apiErr *dto.APIError
	var data dto.UpdateResourceDTO

	resourceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	if err = c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.ID = uint64(resourceID)

	resource, err := usecases.UpdateResource(*h.service, data)
	if err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, resource)
}

func (h *RestRouter) deleteResource(c *gin.Context) {
	var apiErr *dto.APIError
	resourceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "resource ID not specified")
		return
	}

	if err = usecases.DeleteResource(*h.service, uint64(resourceID)); err != nil {
		if errors.As(err, &apiErr) {
			ErrorResponse(c, apiErr.HttpCode, err.Error())
		}
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	EmptyResponse(c)
}
