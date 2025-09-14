package response

const (
	ErrCodeSuccess      = 20001 // Sucess
	ErrCodeParamInvalid = 20003 // Email is invalid
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
}
