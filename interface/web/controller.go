package web

type Contexter interface {
	Param(key string) string
	Query(key string) string
	PostForm(key string) string
	JSON(code int, response interface{})
}

type HandlerFunc = func(context Contexter)
type HandlerFuncs = []HandlerFunc

type Controller interface {
	RegisterRoutes(router Router)
}
