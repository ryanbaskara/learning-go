package config

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type ServerConfig struct {
	ServerHost   string        `envconfig:"SERVER_HOST"`
	ReadTimeout  time.Duration `default:"5m"            envconfig:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `default:"5m"            envconfig:"SERVER_WRITE_TIMEOUT"`

	DatabaseConfig
}

type DatabaseConfig struct {
	Host         string        `default:"127.0.0.1"       envconfig:"DB_MYSQL_HOST"`
	Port         int           `default:"3306"            envconfig:"DB_MYSQL_PORT"`
	Username     string        `default:"learning"        envconfig:"DB_MYSQL_USERNAME"`
	Password     string        `default:"learning"        envconfig:"DB_MYSQL_PASSWORD"`
	Database     string        `default:"learning"        envconfig:"DB_MYSQL_DATABASE"`
	MaxLifetime  time.Duration `default:"4h"              envconfig:"DB_MAX_LIFETIME"`
	MaxIdleTime  time.Duration `default:"5m"              envconfig:"DB_MAX_IDLETIME"`
	MaxIdleConns int           `default:"5"               envconfig:"DB_MAX_IDLECONNS"`
	MaxOpenConns int           `default:"7"               envconfig:"DB_MAX_OPENCONNS"`
}

func loadServerConfig() (ServerConfig, error) {
	var config ServerConfig
	// load from .env if exists
	if _, err := os.Stat(".env"); err == nil {
		if err := gotenv.Load(); err != nil {
			return config, err
		}
	}

	err := envconfig.Process("server", &config)
	return config, err
}

func (c *DatabaseConfig) databaseSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}
