package service

import (
	"dasar/entity"
	"dasar/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// buat category repo dr mock (mock kosong)
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

//repo dr mock
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program terhdp mock. arti dari On = ketika FindById (ada di category_service.go func Get) dipanggil dg parameter (id-nya) adlh 1 maka return nil. artinya
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	// test, data harus kosong
	assert.Nil(t, category)
	// test hrs error, karena data kosong
	assert.NotNil(t, err)
}

// func above error because didnt have data (data is empty). below is success (data is exist)
func TestCategoryService_GetSuccess(t *testing.T) {
	// buat data dummy dulu
	data := entity.Category{
		Id:   "1",
		Name: "jauh",
	}
	// returnnya sdh jd pointer di func FindById category_repository_mock.go
	categoryRepository.Mock.On("FindById", "2").Return(data)
	input, err := categoryService.Get("2")
	// test
	assert.Nil(t, err)                 //error hrs nil karena data ditemukan
	assert.NotNil(t, input)            //data hrs ditemukan / ada
	assert.Equal(t, data.Id, input.Id) // data yg dicari dg data yg dicari lewat id (FindById) harus sama
	assert.Equal(t, data.Name, input.Name)
}
