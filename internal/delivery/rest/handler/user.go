package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martynove/gophermart/internal/models"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		h.logger.Debugf("error parsedd json body: %s", err.Error())
		return
	}
	_, err := h.service.Authorization.CreateUser(input)
	switch {
	case err == nil:
		token, err := h.service.GenerateToken(input.Login, input.Password)
		//c.Status(http.StatusOK)
		if err != nil {
			h.logger.Debugf("GenerateToken %s", err.Error())
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		authHeader := fmt.Sprintf("Bearer %s", token)
		c.Header("Authorization", authHeader)
	case err == models.ErrorLoginExist:
		newErrorResponse(c, http.StatusConflict, err.Error())
		h.logger.Debugf("login: %s already exist", input.Login)
		return
	default:
		newErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		h.logger.Debugf("Internal Server Error: %s", err.Error())
		return
	}
}

func (h *Handler) login(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		h.logger.Debugf("error parsed json body: %s", err.Error())
		return
	}
}
