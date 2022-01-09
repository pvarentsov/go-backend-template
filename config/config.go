package config

import (
	"fmt"
	"time"

	"go-backend-template/api/http"
	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase"
)

// Config

type Config struct {
	httpHost string
	httpPort int

	databaseURL string

	accessTokenExpiresTTL int
	accessTokenSecret     string
}

func TestConfig() *Config {
	return &Config{
		httpHost: "0.0.0.0",
		httpPort: 3000,

		databaseURL: "postgres://go-backend-template:go-backend-template@localhost:5454/go-backend-template",

		accessTokenExpiresTTL: 2 * 60,
		accessTokenSecret:     "secret",
	}
}

func (c *Config) HTTP() http.Config {
	return &httpConfig{
		host: c.httpHost,
		port: c.httpPort,
	}
}

func (c *Config) Usecase() usecase.Config {
	return &usecaseConfig{
		accessTokenExpiresTTL: c.accessTokenExpiresTTL,
		accessTokenSecret:     c.accessTokenSecret,
	}
}

func (c *Config) Database() database.Config {
	return &databaseConfig{
		url: c.databaseURL,
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

func (c *databaseConfig) URL() string {
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
