package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/albertbui010/go-ecommerce-backend-api/internal/repo"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/utils/random"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/albertbui010/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email::::: %s", hashEmail)
	// 5. Check OTP is available

	// 6. User spam email ?

	// 1. Check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. New OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is :::%d\n", otp)

	// 3. Add OTP in redis, set expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. Send Email OTP
	err = sendto.SendTemplateEmailOTP([]string{email}, "buiquangquy12823@gmail.com", "otp-auth.html", map[string]interface{}{"otp": strconv.Itoa(otp)})
	if err != nil {
		return response.ErrSendEmailOTP
	}
	return response.ErrCodeSuccess
}
