package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testProject/internal/entity"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "admin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	admin, err := h.usecases.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, *admin)
}

func (h *Handler) permissionMiddleware(c *gin.Context) {
	var isAdmin string
	user, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user access not found in the context")
		return
	} else if user == false {
		isAdmin = "user"
	} else if user == true {
		isAdmin = "admin"
	}
	permissions, err := h.usecases.RedisUsage.GetPermissionsByRole(isAdmin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var action entity.Premissoins
	action.Url = c.Request.URL.Path

	if !containsPermission(permissions, action) {
		newErrorResponse(c, http.StatusForbidden, "Access is denied")
		return
	}
	c.Next()
}

func containsPermission(permissions []entity.Premissoins, action entity.Premissoins) bool {
	for _, v := range permissions {
		if v == action {
			return true
		}
	}
	return false
}
