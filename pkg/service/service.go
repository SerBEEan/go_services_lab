package repository

import (
	"go_services_lab/pkg/entity"
	"go_services_lab/pkg/repository"
)

type Product interface {
	Create(product entity.Product) (int, error)
	GetAll() ([]entity.Product, error)
	LastOne() (entity.Product, error)
}

type Order interface {
	Get(id int) (entity.Order, error)
	GetAll() ([]entity.Order, error)
	Amount(id int) (float32, error)
	Delete(id int) (int, error)
	Create(user_id int, products map[string]int) (int, error)
}

type Service struct {
	Product
	Order
}

func NewService(rep *repository.Repository) *Service {

	return &Service{
		Product: NewProductService(rep.Product),
		Order:   NewOrderService(rep.Order),
	}
}
