package rest

import (
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestRouter struct {
	service *services.Service
}

func NewRestRouter(service *services.Service) *RestRouter {
	return &RestRouter{service: service}
}

func (h *RestRouter) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	roles := router.Group("/roles", h.JWTAuthMiddleware)
	{
		roles.POST("/", h.createRole)
		roles.GET("/:id", h.retrieveRole)
		roles.GET("/", h.searchRoles)
		roles.PATCH("/", h.updateRole)
		roles.DELETE("/", h.deleteRole)
	}

	return router
}
