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
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		h.logger.Debugf("error created user: %s ", err.Error())
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
