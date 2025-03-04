package config

import (
	"github.com/1995parham/koochooloo/internal/db"
	"github.com/1995parham/koochooloo/internal/logger"
	"github.com/1995parham/koochooloo/internal/metric"
	telemetry "github.com/1995parham/koochooloo/internal/telemetry/config"
)

// Default return default configuration.
func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "debug",
			Syslog: logger.Syslog{
				Enabled: false,
				Network: "",
				Address: "",
				Tag:     "",
			},
		},
		Database: db.Config{
			Name: "koochooloo",
			URL:  "mongodb://127.0.0.1:27017",
		},
		Monitoring: metric.Config{
			Address: ":8080",
			Enabled: true,
		},
		Telemetry: telemetry.Config{
			Trace: telemetry.Trace{
				Enabled: false,
				URL:     "",
			},
		},
	}
}
