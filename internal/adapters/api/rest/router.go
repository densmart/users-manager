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

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signin", h.signIn)
			auth.POST("/refresh", h.refresh)
			auth.POST("/otp", h.checkOtp)
		}
		roles := v1.Group("/roles", h.JWTAuthMiddleware)
		{
			roles.POST("/", h.createRole)
			roles.GET("/:id", h.retrieveRole)
			roles.GET("/", h.searchRoles)
			roles.PATCH("/:id", h.updateRole)
			roles.DELETE("/:id", h.deleteRole)

			permissions := roles.Group(":id/permissions")
			{
				permissions.POST("/", h.createPermissions)
				permissions.GET("/", h.searchPermissions)
				permissions.PATCH("/", h.updatePermissions)
				permissions.DELETE("/", h.deletePermissions)
			}
		}

		resources := v1.Group("/resources", h.JWTAuthMiddleware)
		{
			resources.POST("/", h.createResource)
			resources.GET("/:id", h.retrieveResource)
			resources.GET("/", h.searchResources)
			resources.PATCH("/:id", h.updateResource)
			resources.DELETE("/:id", h.deleteResource)
		}

		users := v1.Group("/users", h.JWTAuthMiddleware)
		{
			users.POST("/", h.createUser)
			users.GET("/:id", h.retrieveUser)
			users.GET("/", h.searchUsers)
			users.PATCH("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}
