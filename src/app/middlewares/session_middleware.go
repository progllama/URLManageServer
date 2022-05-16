package middlewares

import (
	"github.com/gin-gonic/gin"
)

type SessionFactory interface {
	Create(*gin.Context) Session
}

type Session interface {
	HasUserId() bool
}

type LoginRequireMiddleWare struct {
	factory   SessionFactory
	onSuccess func(*gin.Context)
	onFail    func(*gin.Context)
}

func NewLoginRequireMiddleware(factory SessionFactory, onSuccess, onFail func(*gin.Context)) *LoginRequireMiddleWare {
	return &LoginRequireMiddleWare{
		factory,
		onSuccess,
		onFail,
	}
}

func (mw *LoginRequireMiddleWare) Handler() gin.HandlerFunc {
	if mw.factory == nil {
		panic("Factory is nil.")
	}
	if mw.onSuccess == nil || mw.onFail == nil {
		panic(`"onSuccess" and/or "onFail" is nil.`)
	}

	return func(ctx *gin.Context) {
		session := mw.getSession(ctx)
		if session.HasUserId() {
			mw.onSuccess(ctx)
		} else {
			mw.onFail(ctx)
		}
	}
}

func (mw *LoginRequireMiddleWare) getSession(ctx *gin.Context) Session {
	return mw.factory.Create(ctx)
}
