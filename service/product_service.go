package service

import (
	"errors"
	"projek_kesebelas/entity"
	"projek_kesebelas/repository"
	// "go-unit-test/entity"
	// "go-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
