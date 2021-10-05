package serviceimpl

import (
	"go-architecture-mysql/api/entities"
	"go-architecture-mysql/api/payloads"
	"go-architecture-mysql/api/repositories"
	"go-architecture-mysql/api/services"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func CreateUserService(userRepo repositories.UserRepository) services.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (r *UserServiceImpl) ViewUser() ([]entities.User, error) {
	get, err := r.userRepo.View()

	if err != nil {
		return nil, err
	}

	return get, nil
}

func (r *UserServiceImpl) FindUser(username string) (entities.User, error) {
	get, err := r.userRepo.FindByUsername(username)

	if err != nil {
		return entities.User{}, err
	}

	return get, nil
}

func (r *UserServiceImpl) FindById(userId uint) (entities.User, error) {
	get, err := r.userRepo.FindById(userId)

	if err != nil {
		return entities.User{}, err
	}

	return get, nil
}

func (r *UserServiceImpl) UpdateUser(userReq payloads.CreateRequest, userId uint) (bool, error) {
	get, err := r.userRepo.Update(userReq, userId)

	if err != nil {
		return false, err
	}

	return get, nil
}

func (r *UserServiceImpl) DeleteUser(userId uint) (bool, error) {
	get, err := r.userRepo.Delete(userId)

	if err != nil {
		return false, err
	}

	return get, nil
}
