package config

import (
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"os"
	"puppy/models/common/security"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	DefJwtConfig middleware.JWTConfig
}

var AppConfig Config

func GetConfig() error {

	file, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	AppConfig.DefJwtConfig = middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JwtClaims{},
	}
	return nil
}
