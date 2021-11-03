package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/mkaiho/go-oauth-sandbox/interface/web"
)

type ginRouter struct {
	router *gin.Engine
}

func newGinRouter() *ginRouter {
	return &ginRouter{router: gin.Default()}
}

func (r *ginRouter) Get(relativePath string, handlerFunc web.HandlerFunc) {
	r.router.GET(relativePath, func(ctx *gin.Context) {
		ctxWrap := newGinContext(ctx)
		handlerFunc(ctxWrap)
	})
}
