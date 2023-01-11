package rest

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

func (h *RestRouter) createAction(c *gin.Context) {
	var data dto.CreateActionDTO

	if err := c.BindJSON(&data); err != nil {
		ErrorResponse(c, 400, err.Error())
		return
	}

	result, err := usecases.CreateAction(*h.service, data)
	if err != nil {
		ErrorResponse(c, err.HttpCode, err.Error())
		return
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) retrieveAction(c *gin.Context) {
	actionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	result, err := usecases.RetrieveAction(*h.service, uint64(actionID))
	if err != nil {
		ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	SuccessResponse(c, result)
}

func (h *RestRouter) searchActions(c *gin.Context) {
	var data dto.SearchActionDTO

	if err := c.ShouldBindWith(&data, binding.Form); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.RawURL = c.Request.URL.String()

	actions, err := usecases.SearchActions(*h.service, data)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Link", actions.Pagination)
	SuccessResponse(c, actions)
}

func (h *RestRouter) updateAction(c *gin.Context) {
	var data dto.UpdateActionDTO

	actionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "incorrect query string")
		return
	}

	if err = c.BindJSON(&data); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data.ID = uint64(actionID)

	action, err := usecases.UpdateAction(*h.service, data)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SuccessResponse(c, action)
}

func (h *RestRouter) deleteAction(c *gin.Context) {
	actionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "action ID not specified")
		return
	}

	if err = usecases.DeleteAction(*h.service, uint64(actionID)); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	EmptyResponse(c)
}
