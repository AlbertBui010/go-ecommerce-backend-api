package response

const (
	ErrCodeSuccess      = 20001 // Sucess
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 20004
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
}
