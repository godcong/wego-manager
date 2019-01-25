package controller

type Ret struct {
	Code    int    `json:"code" example:"-1"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(msg string) *Ret {
	return &Ret{
		Code:    -1,
		Message: msg,
	}
}
