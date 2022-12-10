package repository

import (
	"go_services_lab/pkg/entity"
	"go_services_lab/pkg/repository"
)

type ProductService struct {
	rep repository.Product
}

func NewProductService(rep repository.Product) *ProductService {
	return &ProductService{rep: rep}
}

func (s *ProductService) Create(product entity.Product) (int, error) {
	return s.rep.Create(product)
}

func (s *ProductService) GetAll() ([]entity.Product, error) {
	return s.rep.GetAll()
}

func (s *ProductService) LastOne() (entity.Product, error) {
	return s.rep.LastOne()
}
