package repository

import (
	"go_services_lab/pkg/entity"

	"github.com/patrickmn/go-cache"
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
	Create(user_id int, products map[int]int) (int, error)
}

type Repository struct {
	Product
	Order
}

func NewRepository(c *cache.Cache) *Repository {
	return &Repository{
		Product: NewProductCache(c),
		Order:   NewOrderCache(c),
	}
}
