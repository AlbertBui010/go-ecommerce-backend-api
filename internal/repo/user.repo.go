package repo

import (
	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/model"
)

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
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}
