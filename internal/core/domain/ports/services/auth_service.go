package services

import "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"

type AuthService interface {
	Login(req entities.StaffLoginRequest) (*entities.StaffLoginRequest, error)
}
