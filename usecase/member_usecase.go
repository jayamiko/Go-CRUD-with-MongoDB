package usecase

import (
	"errors"
	"golang-crud-basic/model"
	"golang-crud-basic/repository"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type MemberUsecase interface {
	GetAll() ([]model.Member, error)
	GetByRecruiterID(id string) (*model.Member, error)
	Create(member *model.Member) error
	UpdateByRecruiter(id string, member *model.Member) error
	Delete(id string) error
}

type memberUsecase struct {
	memberRepo repository.MemberRepository
}

func NewMemberUsecase(memberRepo repository.MemberRepository) MemberUsecase {
	return &memberUsecase{memberRepo}
}

func (uc *memberUsecase) GetAll() ([]model.Member, error) {
	return uc.memberRepo.GetAll()
}

func (uc *memberUsecase) GetByRecruiterID(recruiterID string) (*model.Member, error) {
	return uc.memberRepo.GetByRecruiterID(recruiterID)
}


func (uc *memberUsecase) Create(member *model.Member) error {
    // Generate recruiterId otomatis
    member.RecruiterID = primitive.NewObjectID()

    // Validasi statusAktivasi
    if member.StatusAktivasi != "ACTIVE" && member.StatusAktivasi != "INACTIVE" && member.StatusAktivasi != "PENDING" {
        return errors.New("invalid statusAktivasi, must be ACTIVE, INACTIVE, or PENDING")
    }

    // Validasi email format
    match, _ := regexp.MatchString(`^.+@.+\..+$`, member.Email)
    if !match {
        return errors.New("invalid email format")
    }

    // **Cek duplicate email**
    exists, err := uc.memberRepo.ExistsByEmail(member.Email)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("email already in use")
    }

    // Password wajib dan hash
    if member.Password == "" {
        return errors.New("password is required")
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    member.Password = string(hashedPassword)

    // Set timestamps
    member.CreatedAt = time.Now()
    member.UpdatedAt = time.Now()

    return uc.memberRepo.Create(member)
}

func (uc *memberUsecase) UpdateByRecruiter(recruiterID string, member *model.Member) error {
	// Validasi statusAktivasi
	if member.StatusAktivasi != "ACTIVE" && member.StatusAktivasi != "INACTIVE" && member.StatusAktivasi != "PENDING" {
		return errors.New("invalid statusAktivasi")
	}

	// Validasi email
	match, _ := regexp.MatchString(`^.+@.+\..+$`, member.Email)
	if !match {
		return errors.New("invalid email format")
	}

	// Password wajib
	if member.Password == "" {
		return errors.New("password is required")
	}

	member.UpdatedAt = time.Now()

	return uc.memberRepo.UpdateByRecruiter(recruiterID, member)
}


func (uc *memberUsecase) Delete(id string) error {
	return uc.memberRepo.Delete(id)
}
