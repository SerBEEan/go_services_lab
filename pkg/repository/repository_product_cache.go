package repository

import (
	"errors"
	"go_services_lab/pkg/entity"
	"strconv"

	"github.com/patrickmn/go-cache"
)

type ProductCache struct {
	c *cache.Cache
}

func NewProductCache(c *cache.Cache) *ProductCache {
	return &ProductCache{c: c}
}

func (r *ProductCache) getCount() (int, error) {
	ret, f := r.c.Get("countProduct")
	if !f {
		return 0, errors.New("Unable to get number of products.")
	}
	return ret.(int), nil
}

func (r *ProductCache) getByName(name string) error {
	count, f := r.getCount()
	if f != nil {
		return f
	}
	for i := 1; i <= count; i++ {
		pr, fl := r.c.Get("product" + strconv.Itoa(i))
		if fl {
			if name == pr.(*entity.Product).Name {
				return errors.New("Product with this name exist.")
			}
		}
	}
	return nil
}

func (r *ProductCache) Create(product entity.Product) (int, error) {
	curr_id, f := r.getCount()
	curr_id += 1
	if f != nil {
		return 0, errors.New("Unable to get number of products.")
	}
	fl := r.getByName(product.Name)
	if fl != nil {
		return 0, fl
	}
	r.c.Set("product"+strconv.Itoa(curr_id), &entity.Product{curr_id, product.Name, product.Price}, cache.DefaultExpiration)
	r.c.Increment("countProduct", 1)
	return curr_id, nil
}

func (r *ProductCache) GetAll() ([]entity.Product, error) {
	var retList []entity.Product
	curr_id, f := r.getCount()
	if f != nil {
		return retList, errors.New("Unable to get number of products.")
	}
	for i := 1; i <= curr_id; i++ {
		pr, f := r.c.Get("product" + strconv.Itoa(i))
		if f {
			retList = append(retList, entity.Product{pr.(*entity.Product).ID, pr.(*entity.Product).Name, pr.(*entity.Product).Price})
		}
	}
	return retList, nil
}

func (r *ProductCache) LastOne() (entity.Product, error) {
	var retProduct entity.Product
	curr_id, f := r.getCount()
	if f != nil {
		return retProduct, errors.New("Unable to get number of products.")
	}
	pr, fl := r.c.Get("product" + strconv.Itoa(curr_id))
	if !fl {
		return retProduct, errors.New("Unable to get last one.")
	}
	return entity.Product{pr.(*entity.Product).ID, pr.(*entity.Product).Name, pr.(*entity.Product).Price}, nil
}
