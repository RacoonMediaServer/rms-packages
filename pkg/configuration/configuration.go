package configuration

import (
	"fmt"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
)

// Database is required fields for connecting to database
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Monitor is monitoring settings
type Monitor struct {
	Enabled bool
	Host    string
	Port    int
}

// Debug represents debug settings
type Debug struct {
	Verbose bool
	Monitor Monitor
}

// Http settings of the web server
type Http struct {
	Host string
	Port int
}

// Load read and parses configuration file. Also override some fields via environment variables
// Ex. you can set environment variable RMS_DATABASE_USER which override value root / database / user
func Load(configFile string, configuration interface{}) error {
	err := config.Load(
		file.NewSource(file.WithPath(configFile)),
		env.NewSource(env.WithStrippedPrefix("RMS")),
	)

	if err != nil {
		return err
	}

	if err := config.Scan(&configuration); err != nil {
		return err
	}

	logger.Infof("Configuration: %+v", configuration)

	return nil
}

// GetConnectionString composes PostgresSQL connection string to RMS database
func (config Database) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
}
