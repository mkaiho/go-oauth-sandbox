package main

import (
	"github.com/mkaiho/go-oauth-sandbox/infrastructure"
	"github.com/mkaiho/go-oauth-sandbox/interface/web"
)

func init() {
	web.NewServer = infrastructure.NewGinServer
}
