package errx

type CodeError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func NewCodeError(code uint32, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}
func NewErrCode(code uint32) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  message[code],
	}
}

func NewErrMsg(msg string) *CodeError {
	return &CodeError{
		Code: SERVER_COMMON_ERROR,
		Msg:  msg,
	}
}
