package infrastructure

import "os"

var Env EnvLoader = NewSimpleEnvLoader()

type EnvLoader interface {
	Get(key string) string
}

func NewSimpleEnvLoader() EnvLoader {
	return &simpleEnvLoader{}
}

type simpleEnvLoader struct {
}

func (l *simpleEnvLoader) Get(key string) string {
	return os.Getenv(key)
}
