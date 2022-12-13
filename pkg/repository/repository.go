package repository

import (
	"go_services_lab/models"

	"github.com/patrickmn/go-cache"
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
	Create(user_id int, products map[int]int) (int, error)
}

type OrderRepository struct {
	Product
	Order
}

type User interface {
	Get(id int) (models.User, error)
	Create(models.User) (int, error)
	GetAll() ([]models.User, error)
	Delete(id int) (int, error)
}

type UserRepository struct {
	User
}

func NewRepositoryOrder(c *cache.Cache) *OrderRepository {
	return &OrderRepository{
		Product: NewProductCache(c),
		Order:   NewOrderCache(c),
	}
}

func NewRepositoryUser(c *cache.Cache) *UserRepository {
	return &UserRepository{
		User: NewUserCache(c),
	}
}
