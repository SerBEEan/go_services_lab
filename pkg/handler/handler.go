package handler

import (
	service "go_services_lab/pkg/service"

	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	services *service.ServiceOrder
}

func NewHandlerOrder(services *service.ServiceOrder) *HandlerOrder {
	return &HandlerOrder{services: services}
}

type HandlerUser struct {
	services *service.ServiceUser
}

func NewHandlerUser(services *service.ServiceUser) *HandlerUser {
	return &HandlerUser{services: services}
}

func (h *HandlerOrder) InitRoutesOrder() *gin.Engine {
	router := gin.New()
	product := router.Group("/product")
	{
		product.POST("/all", h.getProductList)
		product.POST("/add", h.addProduct)
		product.POST("/last", h.lastProduct)
	}
	order := router.Group("/order")
	{
		order.POST("/get/:id", h.getOrderById)
		order.POST("/del/:id", h.deleteOrder)
		order.POST("/amount/:id", h.calcAmountOrder)
		order.POST("/add", h.addOrder)
		order.POST("/all", h.getOrderList)
	}

	return router
}

func (h *HandlerUser) InitRoutesUser() *gin.Engine {
	router := gin.New()
	router.POST("/get/:id", h.getUserByID)
	router.POST("/del/:id", h.deleteUser)
	router.POST("/add", h.addUser)
	router.POST("/all", h.getUserList)

	return router
}
