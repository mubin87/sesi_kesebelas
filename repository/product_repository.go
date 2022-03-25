package repository

import "projek_kesebelas/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
