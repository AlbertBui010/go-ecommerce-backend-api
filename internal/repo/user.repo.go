package repo

// type UserRepository struct{}

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{}
// }

// func (ur *UserRepository) GetUserById() string {
// 	return "Hello"
// }

// Interface version
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}
