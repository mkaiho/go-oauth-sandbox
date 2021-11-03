package main

import (
	"github.com/mkaiho/go-oauth-sandbox/interface/web"
)

func main() {
	server := web.NewServer()

	web.RegisterAllControllers(server.Router())

	server.Run("localhost", 3000)
}
