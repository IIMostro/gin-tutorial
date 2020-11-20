package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var InnerException = BizException{Code: 500, Msg: "inner exception!"}

type Exception interface {
	GetCode() int
	GetMsg() string
}

type BizException struct {
	Code int
	Msg  string
}

func (ex BizException) GetCode() int {
	return ex.Code
}

func (ex BizException) GetMsg() string {
	return ex.Msg
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, Failure(errorToString(r)))
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) (int, string) {
	switch v := r.(type) {
	case error:
		return 500, v.Error()
	case Exception:
		return v.GetCode(), v.GetMsg()
	default:
		return 500, r.(string)
	}
}
