package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	//read header
	authHeader := strings.Split(c.GetHeader("Authorization"), " ")
	if len(authHeader) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		c.Abort()
		return
	}
	accessToken := authHeader[1]
	//validate token
	userID, err := h.service.ParseToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		c.Abort()
		return
	}
	h.userID = userID
	c.Next()
}
