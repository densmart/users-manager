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

	actions := router.Group("/actions", h.JWTAuthMiddleware)
	{
		actions.POST("/", h.createAction)
		actions.GET("/:id", h.retrieveAction)
		actions.GET("/", h.searchActions)
		actions.PATCH("/", h.updateAction)
		actions.DELETE("/", h.deleteAction)
	}

	resources := router.Group("/resources", h.JWTAuthMiddleware)
	{
		resources.POST("/", h.createResource)
		resources.GET("/:id", h.retrieveResource)
		resources.GET("/", h.searchResources)
		resources.PATCH("/", h.updateResource)
		resources.DELETE("/", h.deleteResource)
	}

	return router
}
