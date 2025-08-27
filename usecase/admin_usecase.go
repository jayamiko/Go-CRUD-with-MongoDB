package usecase

import (
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
)

type AdminUsecase interface {
	GetAll() ([]model.Admin, error)
	GetByEmail(email string) (*model.Admin, error)
	Create(admin *model.Admin) error
	UpdateByEmail(email string, admin *model.Admin) error
	DeleteByEmail(email string) error
}

type adminUsecase struct {
	adminRepo repository.AdminRepository
}

func NewAdminUsecase(adminRepo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{adminRepo}
}

func (uc *adminUsecase) GetAll() ([]model.Admin, error) {
	return uc.adminRepo.GetAll()
}

func (uc *adminUsecase) GetByEmail(email string) (*model.Admin, error) {
	return uc.adminRepo.GetByEmail(email)
}

func (uc *adminUsecase) Create(admin *model.Admin) error {
	return uc.adminRepo.Create(admin)
}

func (uc *adminUsecase) UpdateByEmail(email string, admin *model.Admin) error {
	return uc.adminRepo.UpdateByEmail(email, admin)
}

func (uc *adminUsecase) DeleteByEmail(email string) error {
	return uc.adminRepo.DeleteByEmail(email)
}
