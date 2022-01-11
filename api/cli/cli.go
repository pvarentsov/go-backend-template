package cli

import (
	"github.com/alecthomas/kong"
	"go-backend-template/config"
)

type CLI struct {
	envPath string `help:"(Optional) Path to env config file" optional:""`
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) ParseConfig() (*config.Config, error) {
	kong.Parse(&c)

	return config.ParseEnv(c.envPath)
}
