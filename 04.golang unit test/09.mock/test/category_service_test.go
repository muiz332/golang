package test

import (
	"mock/entity"
	"mock/repository"
	"mock/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

var categoryService = service.CategoryService{
	Repository: categoryRepository,
}

func TestCategoryService_GetNotFound(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")

	assert.NotNil(t, err)
	assert.Nil(t, category)

}
func TestCategoryService_GetFound(t *testing.T) {

	var data entity.Category = entity.Category{
		Id:   "2",
		Name: "hasan",
	}

	categoryRepository.Mock.On("FindById", "2").Return(data)
	category, err := categoryService.Get("2")

	assert.NotNil(t, category)
	assert.Nil(t, err)

}
