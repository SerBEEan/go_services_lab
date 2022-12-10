package handler

import (
	"go_services_lab/pkg/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addProduct(c *gin.Context) {
	var input entity.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body\n")
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProductList(c *gin.Context) {
	productList, err := h.services.Product.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range productList {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":    productList[i].ID,
			"name":  productList[i].Name,
			"price": productList[i].Price,
		})
	}
}

func (h *Handler) lastProduct(c *gin.Context) {
	product, err := h.services.Product.LastOne()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    product.ID,
		"name":  product.Name,
		"price": product.Price,
	})
}
