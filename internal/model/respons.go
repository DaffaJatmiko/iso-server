package model

type SuccessResp struct {
	Message string `json:"message"`
}

type ErrorResp struct {
	Message string `json:"message"`
}

func NewSuccessResp(message string) *SuccessResp {
	return &SuccessResp{message}
}

func NewErrorResp(message string) *ErrorResp {
	return &ErrorResp{message}
}
