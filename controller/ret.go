package controller

type Ret struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(msg string) *Ret {
	return &Ret{
		Code:    -1,
		Message: msg,
	}
}
