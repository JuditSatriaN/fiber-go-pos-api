package server

import (
	"os"

	"github.com/fiber-go-pos-api/internal/app/constant"
)

func GetAppPort() string {
	port := os.Getenv(constant.AppPortEnvKey)
	if port == "" {
		return constant.DefaultPort
	}
	return port
}
