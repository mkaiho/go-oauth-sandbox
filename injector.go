package main

import (
	"github.com/mkaiho/go-oauth-sandbox/infra"
	"github.com/mkaiho/go-oauth-sandbox/interface/web"
)

func init() {
	web.NewServer = infra.NewGinServer
}
