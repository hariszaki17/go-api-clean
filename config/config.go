package config

import (
	"github.com/hariszaki17/go-api-clean/exception"
	"os"
	"github.com/joho/godotenv"
)
// Config expose global
type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

// New expose global
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}