package serviceimpl

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
	"go-architecture-mysql/api/repositories"
	"go-architecture-mysql/api/services"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func CreateAuthService(userRepo repositories.UserRepository) services.AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) DoLogin(userReq payloads.LoginRequest) (entities.User, error) {
	get, err := s.userRepo.FindByUsername(userReq.Username)

	if err != nil {
		return entities.User{}, err
	}

	return get, nil
}

func (s *AuthServiceImpl) DoRegister(userReq payloads.CreateRequest) (bool, error) {
	get, err := s.userRepo.Create(userReq)

	if err != nil {
		return false, err
	}

	return get, nil
}