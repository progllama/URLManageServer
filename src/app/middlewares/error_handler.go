package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	Handle(*gin.Context)
}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}

// フレームワーク（gin)によるエラーのハンドリングのみ行う。
type errorHandler struct {
}

func (handler *errorHandler) Handle(ctx *gin.Context) {
	ctx.Next()

	errs := ctx.Errors
	if len(errs) <= 0 {
		return
	}

	err := errs.Last()
	code := http.StatusInternalServerError
	message := "unknown error happen"
	switch err.Type {
	case gin.ErrorTypeRender:
		message = err.Error()
	case gin.ErrorTypeBind:
		message = err.Error()
	case gin.ErrorTypePrivate:
		message = err.Error()
	case gin.ErrorTypePublic:
		message = err.Error()
	case gin.ErrorTypeAny:
		message = err.Error()
	}

	ctx.JSON(code, message)
	ctx.Abort()
}
