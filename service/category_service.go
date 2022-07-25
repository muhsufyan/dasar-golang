package service

import (
	"dasar/entity"
	"dasar/repository"
	"errors"
)

// idealnya service ini dibuat juga interface-nya tp ditutorial tdk dibuat
// https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41
// https://threedots.tech/post/repository-pattern-in-go/
type CategoryService struct {
	Repository repository.ICategoryRepository
}

func (catserv CategoryService) Get(id string) (*entity.Category, error) {
	category := catserv.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
