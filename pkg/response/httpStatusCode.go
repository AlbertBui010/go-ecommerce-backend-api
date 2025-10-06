package response

const (
	ErrCodeSuccess      = 20001 // Sucess
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001
	ErrInvalidOTP   = 30002
	ErrSendEmailOTP = 30003

	// Register
	ErrCodeUserHasExists = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
	ErrInvalidOTP:       "OTP error",
	ErrSendEmailOTP:     "Failed to send email OTP",

	ErrCodeUserHasExists: "User has already registered",
}
