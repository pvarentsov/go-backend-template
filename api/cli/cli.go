package cli

import (
	"github.com/alecthomas/kong"
	"go-backend-template/config"
)

type CLI struct {
	EnvPath string `help:"Path to env config file" type:"path" optional:""`
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) ParseConfig() (*config.Config, error) {
	kong.Parse(c)

	return config.ParseEnv(c.EnvPath)
}
