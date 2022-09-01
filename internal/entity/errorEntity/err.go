package errorEntity

type ErrorRes struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func NewErrorRes(errorCode int, errorMsg string) *ErrorRes {
	return &ErrorRes{ErrorCode: errorCode, ErrorMsg: errorMsg}
}
