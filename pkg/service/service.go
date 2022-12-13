package repository

import (
	"go_services_lab/models"
	"go_services_lab/pkg/repository"
)

type Product interface {
	Create(product models.Product) (int, error)
	GetAll() ([]models.Product, error)
	LastOne() (models.Product, error)
}

type Order interface {
	Get(id int) (models.Order, error)
	GetAll() ([]models.Order, error)
	Amount(id int) (float32, error)
	Delete(id int) (int, error)
	Create(user_id int, products map[string]int) (int, error)
}

type ServiceOrder struct {
	Product
	Order
}

type User interface {
	Get(id int) (models.User, error)
	Create(models.User) (int, error)
	GetAll() ([]models.User, error)
	Delete(id int) (int, error)
}

type ServiceUser struct {
	User
}

func NewServiceOrder(rep *repository.OrderRepository) *ServiceOrder {
	return &ServiceOrder{
		Product: NewProductService(rep.Product),
		Order:   NewOrderService(rep.Order),
	}
}
func NewServiceUser(rep *repository.UserRepository) *ServiceUser {
	return &ServiceUser{
		User: NewUserService(rep.User),
	}
}
