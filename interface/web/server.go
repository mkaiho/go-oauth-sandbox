package web

var NewServer func() Server

type Server interface {
	Router() Router
	Run(host string, port int) error
}
