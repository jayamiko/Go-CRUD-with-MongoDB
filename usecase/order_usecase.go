package usecase

import (
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
)

type OrderUsecase interface {
    Create(order *model.Order) error
    GetAll() ([]model.Order, error)
    GetByID(id string) (*model.Order, error)
}

type orderUsecase struct {
    orderRepo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) OrderUsecase {
    return &orderUsecase{orderRepo: repo}
}

func (uc *orderUsecase) Create(order *model.Order) error {
    return uc.orderRepo.Create(order)
}

func (uc *orderUsecase) GetAll() ([]model.Order, error) {
    return uc.orderRepo.GetAll()
}

func (uc *orderUsecase) GetByID(id string) (*model.Order, error) {
    return uc.orderRepo.GetByID(id)
}
