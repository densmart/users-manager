package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *RestRouter) JWTAuthMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		ErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		ErrorResponse(c, http.StatusUnauthorized, "incorrect authorization header")
		return
	}
}
