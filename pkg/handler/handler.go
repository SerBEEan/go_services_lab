package handler

import (
	service "go_services_lab/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	product := router.Group("/product")
	{
		product.POST("/all", h.getProductList)
		product.POST("/add", h.addProduct)
		product.POST("/last", h.lastProduct)
	}
	order := router.Group("/order")
	{
		order.POST("/get", h.getOrderById)
		order.POST("/del", h.deleteOrder)
		order.POST("/amount", h.calcAmountOrder)
		order.POST("/add", h.addOrder)
		order.POST("/all", h.getOrderList)
	}

	return router
}
