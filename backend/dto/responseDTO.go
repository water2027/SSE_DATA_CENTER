package dto

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Option func(*response)

func WithCode(code int) Option {
	return func(r *response) {
		r.Code = code
	}
}

func WithMessage(msg string) Option {
	return func(r *response) {
		r.Message = msg
	}
}

func WithData(data interface{}) Option {
	return func(r *response) {
		r.Data = data
	}
}

func responseDTO(c *gin.Context, res response) {
	c.JSON(200, res)
}

func Success(c *gin.Context, opts ...Option) {
	res := response{
		Code:    1,
		Message: "成功",
	}
	for _, opt := range opts {
		opt(&res)
	}
	responseDTO(c, res)
}

func Fail(c *gin.Context, opts ...Option) {
	res := response{
		Code:    0,
		Message: "失败",
		Data:    nil,
	}
	for _, opt := range opts {
		opt(&res)
	}
	responseDTO(c, res)
}
