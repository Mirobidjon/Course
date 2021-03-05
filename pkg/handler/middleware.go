package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h Handler) identify(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPart := strings.Split(header, " ")
	if len(headerPart) != 2 {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	role, id, err := h.service.AuthMasters.ParseToken(headerPart[1])
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("id", id)
	c.Set("role", role)
}

func (h Handler) getRoleAndID(c *gin.Context) (string, int, error) {
	id, ok1 := c.Get("id")
	role, ok2 := c.Get("role")
	if !ok1 || !ok2 {
		NewErrorResponce(c, http.StatusInternalServerError, "id or role not found")
		return "", 0, errors.New("id or role not found")
	}

	idInt, ok1 := id.(int)
	roleString, ok2 := role.(string)
	if !ok1 || !ok2 {
		NewErrorResponce(c, http.StatusInternalServerError, "id or role not found")
		return "", 0, errors.New("id or role not found")
	}

	return roleString, idInt, nil
}
