package web

type Method string

const (
	GetMethod Method = "Get"
)

type Route struct {
	Method       Method
	RelativePath string
	HandlerFunc  HandlerFunc
}

type Routes = []Route

type Router interface {
	Get(relativePath string, handlerFunc HandlerFunc)
}

func RegisterAllControllers(router Router) {
	NewUserController().RegisterRoutes(router)
}
