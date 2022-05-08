package handler

import (
	"github.com/gin-gonic/gin"
	"log"
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
	_, err := h.service.ParseToken(accessToken)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		c.Abort()
		return
	}
	c.Next()
}
