package usecase

import (
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
)


type ProductUsecase interface {
    Create(product *model.Product) error
    GetAll() ([]model.Product, error)
    GetByID(id string) (*model.Product, error)
    Update(id string, product *model.Product) error
    Delete(id string) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo}
}

func (uc *productUsecase) Create(product *model.Product) error {
    return uc.productRepo.Create(product)
}

func (uc *productUsecase) GetAll() ([]model.Product, error) {
    return uc.productRepo.GetAll()
}

func (uc *productUsecase) GetByID(id string) (*model.Product, error) {
    return uc.productRepo.GetByID(id)
}

func (uc *productUsecase) Update(id string, product *model.Product) error {
    return uc.productRepo.Update(id, product)
}

func (uc *productUsecase) Delete(id string) error {
    return uc.productRepo.Delete(id)
}
