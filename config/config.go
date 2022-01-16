package config

import (
	"fmt"
	"time"

	"go-backend-template/api/http"
	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

// Config

type Config struct {
	HttpHost string `envconfig:"HTTP_HOST"`
	HttpPort int    `envconfig:"HTTP_PORT"`

	DatabaseURL string `envconfig:"DATABASE_URL"`

	AccessTokenExpiresTTL int    `envconfig:"ACCESS_TOKEN_EXPIRES_TTL"`
	AccessTokenSecret     string `envconfig:"ACCESS_TOKEN_SECRET"`
}

func ParseEnv(envPath string) (*Config, error) {
	if envPath != "" {
		if err := gotenv.OverLoad(envPath); err != nil {
			return nil, err
		}
	}

	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) HTTP() http.Config {
	return &httpConfig{
		host: c.HttpHost,
		port: c.HttpPort,
	}
}

func (c *Config) Usecase() usecase.Config {
	return &usecaseConfig{
		accessTokenExpiresTTL: c.AccessTokenExpiresTTL,
		accessTokenSecret:     c.AccessTokenSecret,
	}
}

func (c *Config) Database() database.Config {
	return &databaseConfig{
		url: c.DatabaseURL,
	}
}

// HTTP

type httpConfig struct {
	host string
	port int
}

func (c *httpConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

// Database

type databaseConfig struct {
	url string
}

func (c *databaseConfig) ConnString() string {
	return c.url
}

// Usecase

type usecaseConfig struct {
	accessTokenExpiresTTL int
	accessTokenSecret     string
}

func (c *usecaseConfig) AccessTokenSecret() string {
	return c.accessTokenSecret
}

func (c *usecaseConfig) AccessTokenExpiresDate() time.Time {
	duration := time.Duration(c.accessTokenExpiresTTL)
	return time.Now().UTC().Add(time.Minute * duration)
}
