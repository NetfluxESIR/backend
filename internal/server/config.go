package server

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

var (
	ErrInvalidPort   = errors.New("invalid port")
	ErrInvalidHost   = errors.New("invalid host")
	ErrInvalidDSN    = errors.New("invalid DSN")
	ErrInvalidLogger = errors.New("invalid logger")
)

type Config struct {
	// Port is the port to listen on.
	Port int
	// Host is the host to listen on.
	Host string
	// DSN is the data source name.
	DSN string
	// Logger is the logger entry.
	Logger *log.Entry
	// AdminAccount is the admin account.
	AdminAccount string
	// AdminPassword is the admin password.
	AdminPassword string
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.Port <= 0 {
		return ErrInvalidPort
	}
	if c.Host == "" {
		return ErrInvalidHost
	}
	if c.DSN == "" {
		return ErrInvalidDSN
	}
	if c.Logger == nil {
		return ErrInvalidLogger
	}
	return nil
}
