package repository

import (
	"dasar/entity"

	"github.com/stretchr/testify/mock"
)

// ini akan implement interface dari ICategoryRepository
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// impement fungsi FindById dg mock
func (repo *CategoryRepositoryMock) FindById(id string) *entity.Category {
	// buat program untuk mock select *
	// calling mock, paramnya adlh param dari fungsi ini dan akan return array (isi semua data)
	arguments := repo.Mock.Called(id)
	// 0 artinya data ke 1 (index ke 0)
	if arguments.Get(0) == nil {
		return nil
	}
	// tdk null, maka data ditangkap oleh struct
	data := arguments.Get(0).(entity.Category)
	// data ini pointer jd perlu tambah &
	return &data
}
