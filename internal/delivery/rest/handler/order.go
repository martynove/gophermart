package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/martynove/gophermart/internal/models"
	"github.com/theplant/luhn"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) getAllUploadOrders(c *gin.Context) {
	var orders []models.Order
	orders, err := h.service.GetAllOrders(h.userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Debugf(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &orders)
}

func (h *Handler) uploadOrder(c *gin.Context) {
	if !strings.Contains(c.ContentType(), "text/plain") {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	number, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || string(number) == "" {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// validate order number by luhn method
	orderNumberInt, err := strconv.Atoi(string(number))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	if !luhn.Valid(orderNumberInt) {
		c.String(http.StatusUnprocessableEntity, "Bad format for order number")
		return
	}
	if err := h.service.UploadOrder(h.userID, orderNumberInt); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
