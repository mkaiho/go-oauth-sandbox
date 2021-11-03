package infra

import (
	"fmt"

	"github.com/mkaiho/go-oauth-sandbox/interface/web"
)

type ginServer struct {
	router *ginRouter
}

func NewGinServer() web.Server {
	return &ginServer{router: newGinRouter()}
}

func (s *ginServer) Run(host string, port int) error {
	return s.router.router.Run(fmt.Sprintf("%s:%d", host, port))
}

func (s *ginServer) Router() web.Router {
	return s.router
}
