package res

import (
	"backend-go/utils/valid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code int

const (
	SUCCESS   Code = 0
	ErrCode   Code = 7 // 系统错误
	validCode Code = 9 // 校验错误
)

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}

func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Data: data,
		Msg:  msg,
	})
}

func OKWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	OK(data, "成功", c)
}

func OKWithList[T any](list []T, count int, c *gin.Context) {
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List:  list,
		Count: count,
	}, "成功", c)
}

func Fail(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(ErrCode, map[string]any{}, msg, c)
}

func FailWithValidMsg(msg string, c *gin.Context) {
	Fail(validCode, map[string]any{}, msg, c)
}

func FailWithData(data any, c *gin.Context) {
	Fail(ErrCode, data, "系统错误", c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	errorMsg := valid.Error(err)
	FailWithMsg(errorMsg, c)
}
