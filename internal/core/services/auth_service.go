package services

import (
	"errors"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type AuthServiceImpl struct {
	staffRepo repositories.StaffRepository
}

func NewAuthService(staffRepo repositories.StaffRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		staffRepo: staffRepo,
	}
}

func (s *AuthServiceImpl) Login(req entities.StaffLoginRequest) (*entities.StaffLoginResponse, error) {
	staff, err := s.staffRepo.GetStaffByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	token, err := utils.GenerateJWT(req.Email)
	if err != nil {
		return nil, errors.New("failed to genterate token")
	}

	return &entities.StaffLoginResponse{
		Token: token,
		Staff: *staff,
	}, nil
}
