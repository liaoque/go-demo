package exception

import (
	"net/http"
)

type ExceptionMzq struct {
	HttpCode int
	Code int	`json:"code,int"`
	Mes string	`json:"mes,string"`
}


const (
	ERROR_PASSWORD = -1
	ERROR_NOT_FOUNT = -2
	ERROR_CAPTCHA = -2
	ERROR_NO_USER = -2
	ERROR_CREATE_TOKEN = -1000
	ERROR_UPDATE = -1000
)

func NewExceptionMzq(code int,mes string, options... int) *ExceptionMzq {
	httpCode := http.StatusUnauthorized
	if len(options) > 0{
		httpCode = options[0]
	}
	return &ExceptionMzq{httpCode,code, mes}
}


func (e *ExceptionMzq) Error() string {

	return e.Mes
}