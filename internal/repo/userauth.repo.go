package repo

import (
	"fmt"
	"time"

	"github.com/albertbui010/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email)
	fmt.Printf("KEY:::%s", key)
	return global.Rdb.SetEX(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
