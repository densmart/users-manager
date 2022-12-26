package rest

import (
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestRouter struct {
	services *services.Service
}

func NewRestRouter(services *services.Service) *RestRouter {
	return &RestRouter{services: services}
}

func (h *RestRouter) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	return router
}
