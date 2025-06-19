package util

import (
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kelseyhightower/envconfig"
)

type CORSConfig struct {
	AllowOrigins []string `json:"allow_origins"  envconfig:"ALLOWED_ORIGINS"`
}

func GenerateCORSConfig() (*cors.Config, error) {
	var cc CORSConfig
	if err := envconfig.Process(ENVCONFIG_PREFIX, &cc); err != nil {
		return nil, err
	}
	return &cors.Config{
		AllowOrigins:     strings.Join(cc.AllowOrigins, ","),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}, nil
}
