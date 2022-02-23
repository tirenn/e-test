package auth

// import (
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	"kata.ai/chatcommerce/inventory-service/models"
// 	"kata.ai/chatcommerce/inventory-service/transport/http/messages/request"
// )

// func Init() (*RepositoryMock, Service) {
// 	var repo = &RepositoryMock{Mock: mock.Mock{}}
// 	return repo, Service{repository: repo}
// }

// func TestCreateSuccess(t *testing.T) {
// 	repo, service := Init()

// 	req := request.Category{
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-1",
// 		Content:    "",
// 		MerchantID: "merchant-1",
// 	}

// 	category := models.Category{
// 		BaseModel: models.BaseModel{
// 			ID: "category-id-1",
// 		},
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-1",
// 		Content:    "",
// 		Slug:       "title-1",
// 		MerchantID: "merchant-1",
// 	}

// 	repo.Mock.On("Create", mock.AnythingOfType("*models.Category")).Return(nil).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*models.Category)
// 		*arg = category
// 	})

// 	res := service.Create(req)
// 	assert.Equal(t, nil, res)
// }

// func TestCreateFail(t *testing.T) {
// 	repo, service := Init()

// 	req := request.Category{
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-1",
// 		Content:    "",
// 		MerchantID: "merchant-1",
// 	}

// 	err := errors.New("Cannot Create")
// 	repo.Mock.On("Create", mock.AnythingOfType("*models.Category")).Return(err)

// 	res := service.Create(req)
// 	assert.Equal(t, err, res)
// }

// func TestGetSuccess(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-1"

// 	category := models.Category{
// 		BaseModel: models.BaseModel{
// 			ID: "category-id-1",
// 		},
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-meta-1",
// 		Content:    "",
// 		Slug:       "title-1",
// 		MerchantID: "merchant-1",
// 	}

// 	repo.Mock.On("GetById", id).Return(category, nil)

// 	res, err := service.Get(id)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, res, category)
// }

// func TestGetFail(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-1"
// 	category := models.Category{}

// 	errGet := errors.New("Not Found")
// 	repo.Mock.On("GetById", id).Return(category, errGet)

// 	res, err := service.Get(id)
// 	assert.Equal(t, err, errGet)
// 	assert.Equal(t, models.Category{}, res)
// }

// func TestUpdateSuccess(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-1"

// 	req := request.Category{
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-1",
// 		Content:    "",
// 		MerchantID: "merchant-1",
// 	}

// 	category := models.Category{
// 		BaseModel: models.BaseModel{
// 			ID: "category-id-1",
// 		},
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-meta-1",
// 		Content:    "",
// 		Slug:       "title-1",
// 		MerchantID: "merchant-1",
// 	}

// 	repo.Mock.On("GetById", id).Return(category, nil)

// 	repo.Mock.On("Update", mock.AnythingOfType("*models.Category")).Return(nil)

// 	err := service.Update(id, req)
// 	assert.Equal(t, nil, err)
// }

// func TestUpdateFail(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-1"

// 	category := models.Category{
// 		BaseModel: models.BaseModel{
// 			ID: "category-id-1",
// 		},
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-meta-1",
// 		Content:    "",
// 		Slug:       "title-1",
// 		MerchantID: "merchant-1",
// 	}

// 	repo.Mock.On("GetById", id).Return(category, nil)

// 	req := request.Category{
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-1",
// 		Content:    "",
// 		MerchantID: "merchant-1",
// 	}

// 	errUpdate := errors.New("Cannot Update")
// 	repo.Mock.On("Update", mock.AnythingOfType("*models.Category")).Return(errUpdate)

// 	err := service.Update(id, req)
// 	assert.Equal(t, errUpdate, err)
// }

// func TestDeleteSuccess(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-id-1"

// 	category := models.Category{
// 		BaseModel: models.BaseModel{
// 			ID: "category-id-1",
// 		},
// 		ParentID:   "parent-1",
// 		Title:      "title 1",
// 		MetaTitle:  "title-meta-1",
// 		Content:    "",
// 		Slug:       "title-1",
// 		MerchantID: "merchant-1",
// 	}

// 	repo.Mock.On("GetById", id).Return(category, nil)

// 	repo.Mock.On("Delete", mock.AnythingOfType("*models.Category")).Return(nil)

// 	err := service.Delete(id)
// 	assert.Equal(t, nil, err)
// }

// func TestDeleteFail(t *testing.T) {
// 	repo, service := Init()

// 	id := "category-1"
// 	category := models.Category{}

// 	repo.Mock.On("GetById", id).Return(category, nil)

// 	errDelete := errors.New("Cannot Delete")
// 	repo.Mock.On("Delete", mock.AnythingOfType("*models.Category")).Return(errDelete)

// 	err := service.Delete(id)
// 	assert.Equal(t, errDelete, err)
// }

// func TestFindSuccess(t *testing.T) {
// 	repo, service := Init()

// 	req := request.FindCategory{
// 		Limit:      10,
// 		Page:       1,
// 		Offset:     0,
// 		Order:      "created_at",
// 		Title:      "title",
// 		MerchantID: "merchant-1",
// 	}

// 	categories := []models.Category{
// 		{
// 			BaseModel: models.BaseModel{
// 				ID: "category-id-1",
// 			},
// 			ParentID:   "parent-1",
// 			Title:      "title 1",
// 			MetaTitle:  "title-meta-1",
// 			Content:    "",
// 			Slug:       "title-1",
// 			MerchantID: "merchant-1",
// 		},
// 		{
// 			BaseModel: models.BaseModel{
// 				ID: "category-id-2",
// 			},
// 			ParentID:   "parent-2",
// 			Title:      "title 2",
// 			MetaTitle:  "title-meta-2",
// 			Content:    "",
// 			Slug:       "title-2",
// 			MerchantID: "merchant-2",
// 		},
// 	}

// 	repo.Mock.On("Find", req).Return(categories, nil)

// 	var countRes int64 = int64(len(categories))
// 	repo.Mock.On("Count", req).Return(countRes, nil)

// 	res, count, err := service.Find(req)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, res, categories)
// 	assert.Equal(t, countRes, count)
// }

// func TestFindNil(t *testing.T) {
// 	repo, service := Init()

// 	req := request.FindCategory{
// 		Limit:      10,
// 		Page:       1,
// 		Offset:     0,
// 		Order:      "created_at",
// 		Title:      "title",
// 		MerchantID: "merchant-1",
// 	}

// 	categories := []models.Category{}

// 	repo.Mock.On("Find", req).Return(categories, nil)

// 	var countRes int64 = int64(len(categories))
// 	repo.Mock.On("Count", req).Return(countRes, nil)

// 	res, count, err := service.Find(req)
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, res, categories)
// 	assert.Equal(t, countRes, count)
// }

// func TestFindFail(t *testing.T) {
// 	repo, service := Init()

// 	req := request.FindCategory{
// 		Limit:      10,
// 		Page:       1,
// 		Offset:     0,
// 		Order:      "created_at",
// 		Title:      "title",
// 		MerchantID: "merchant-1",
// 	}

// 	categories := []models.Category{}

// 	errFind := errors.New("Cannot Find")
// 	repo.Mock.On("Find", req).Return(categories, errFind)

// 	var countRes int64 = int64(len(categories))
// 	repo.Mock.On("Count", req).Return(countRes, nil)

// 	res, count, err := service.Find(req)
// 	assert.Equal(t, errFind, err)
// 	assert.Equal(t, res, categories)
// 	assert.Equal(t, countRes, count)
// }
