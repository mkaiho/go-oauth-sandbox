package infra

import (
	"github.com/gin-gonic/gin"
)

type ginContext struct {
	ctx *gin.Context
}

func newGinContext(ctx *gin.Context) *ginContext {
	return &ginContext{ctx: ctx}
}

func (ctx *ginContext) Param(key string) string {
	return ctx.ctx.Param(key)
}

func (ctx *ginContext) Query(key string) string {
	return ctx.ctx.Query(key)
}

func (ctx *ginContext) PostForm(key string) string {
	return ctx.ctx.PostForm(key)
}

func (ctx *ginContext) JSON(code int, obj interface{}) {
	ctx.ctx.JSON(code, obj)
}
