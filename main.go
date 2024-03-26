package main

import (
	"log/slog"
	"time"

	"github.com/asdine/storm"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type config struct {
	LeaseTime           time.Duration `default:"10m"`
	Database            string        `default:"./dhpresence.db"`
	Webhook             string
	ReconnectAcceptList []string `default:""`
	NewBlockList        []string `default:""`
}

func main() {
	logger := logrus.WithField("source", "dhcpresence")

	var cfg config
	err := envconfig.Process("dh", &cfg)
	if err != nil {
		logger.WithError(err).Fatal("Failed to parse env")
	}

	slog.Info("Starting dhcpresence")
	devices, err := getDevices()
	if err != nil {
		slog.Error("Failed to get cookie", "error", err)
		return
	}

	for _, device := range devices {
		if device.Active {
			slog.Info("Device found", "device", device.Name, "active", device.Active)
		}
	}

	return

	db, err := storm.Open(cfg.Database)
	defer func() {
		err := db.Set("timestamps", "last_run", time.Now())
		if err != nil {
			logger.WithError(err).Error("Failed to store last run")
		}

		db.Close()
	}()

	if err != nil {
		logger.WithError(err).WithField("path", cfg.Database).Fatal("could not open database")
	}

}
