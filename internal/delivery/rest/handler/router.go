package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/martynove/gophermart/internal/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	logger  *logrus.Logger
	userID  int
}

func NewHandler(logger *logrus.Logger, service *service.Service) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/api/user")
	{
		user.POST("/register", h.register)
		user.POST("/login", h.login)
	}
	orders := router.Group("/api/user", h.AuthMiddleware)
	{
		orders.POST("/orders", h.uploadOrder)
		orders.GET("/orders", h.getAllUploadOrders)
	}
	balance := router.Group("/api/user/balance", h.AuthMiddleware)
	{
		balance.GET("", h.getCurrentBalance)
		balance.POST("/withdraw", h.requestPaymentWithPoints)
		balance.GET("/withdrawals", h.getHistoryPayments)
	}
	return router
}
