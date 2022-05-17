package middlewares

import (
	"net/http"
	"url_manager/app/session"

	"github.com/gin-gonic/gin"
)

type SessionFactory interface {
	Create(*gin.Context) Session
}

type Session interface {
	HasUserId() bool
}

type LoginRequireMiddleWare struct {
	factory          session.SessionFactory
	onUserLoggdIn    func(*gin.Context)
	onUserNotLoggdIn func(*gin.Context)
}

func NewLoginRequireMiddleware(factory session.SessionFactory, onUserLoggdIn, onUserNotLoggdIn func(*gin.Context)) *LoginRequireMiddleWare {
	return &LoginRequireMiddleWare{
		factory,
		onUserLoggdIn,
		onUserNotLoggdIn,
	}
}

func (mw *LoginRequireMiddleWare) RequireLogin() gin.HandlerFunc {
	if mw.factory == nil {
		panic("Session factory is nil.")
	}
	if mw.onUserLoggdIn == nil || mw.onUserNotLoggdIn == nil {
		panic(`Callbacks "onSuccess" and/or "onFail" is nil.`)
	}

	return func(ctx *gin.Context) {
		if mw.logsin(ctx) {
			mw.onUserLoggdIn(ctx)
		} else {
			mw.onUserNotLoggdIn(ctx)
		}
	}
}

func (mw *LoginRequireMiddleWare) logsin(ctx *gin.Context) bool {
	session := mw.getSession(ctx)
	return session.HasUserId()
}

func (mw *LoginRequireMiddleWare) getSession(ctx *gin.Context) Session {
	return mw.factory.Create(ctx)
}

func DoNothing(_ *gin.Context) {
	// 成功時のコールバック、 type Any interface{}みたいなもの。
}

func RedirectToLoginPage(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/login") // TODO remove hard coded.
	ctx.Abort()
}
