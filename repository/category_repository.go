package repository

import "dasar/entity"

type ICategoryRepository interface {
	FindById(id string) *entity.Category
}
