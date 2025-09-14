package service

import "github.com/albertbui010/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepository(),
	}
}

func (us *UserService) GetUserById() string {
	return us.userRepo.GetUserById()
}
