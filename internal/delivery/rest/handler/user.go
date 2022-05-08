package handler

import (
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
		c.Status(http.StatusOK)
		return
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
