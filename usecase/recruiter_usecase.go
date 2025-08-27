package usecase

import (
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
)

type RecruiterUsecase interface {
	Create(recruiter *model.Recruiter) error
	GetAll() ([]model.Recruiter, error)
	GetByID(id string) (*model.Recruiter, error)
	Update(id string, recruiter *model.Recruiter) error
	Delete(id string) error
}

type recruiterUsecase struct {
	repo repository.RecruiterRepository
}

func NewRecruiterUsecase(repo repository.RecruiterRepository) RecruiterUsecase {
	return &recruiterUsecase{repo: repo}
}

func (uc *recruiterUsecase) Create(recruiter *model.Recruiter) error {
	return uc.repo.Create(recruiter)
}

func (uc *recruiterUsecase) GetAll() ([]model.Recruiter, error) {
	return uc.repo.GetAll()
}

func (uc *recruiterUsecase) GetByID(id string) (*model.Recruiter, error) {
	return uc.repo.GetByID(id)
}

func (uc *recruiterUsecase) Update(id string, recruiter *model.Recruiter) error {
	return uc.repo.Update(id, recruiter)
}

func (uc *recruiterUsecase) Delete(id string) error {
	return uc.repo.Delete(id)
}
