package usecase

import (
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
)

type InvoiceUsecase interface {
	Create(invoice *model.Invoice) error
	GetAll() ([]model.Invoice, error)
	GetByID(id string) (*model.Invoice, error)
	Delete(id string) error
}

type invoiceUsecase struct {
	repo repository.InvoiceRepository
}

func NewInvoiceUsecase(repo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo: repo}
}

func (uc *invoiceUsecase) Create(invoice *model.Invoice) error {
	return uc.repo.Create(invoice)
}

func (uc *invoiceUsecase) GetAll() ([]model.Invoice, error) {
	return uc.repo.GetAll()
}

func (uc *invoiceUsecase) GetByID(id string) (*model.Invoice, error) {
	return uc.repo.GetByID(id)
}

func (uc *invoiceUsecase) Delete(id string) error {
	return uc.repo.Delete(id)
}
