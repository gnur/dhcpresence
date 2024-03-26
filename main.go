package main

import (
	"log/slog"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LeaseTime           time.Duration `default:"10m"`
	Database            string        `default:"./dhpresence.db"`
	Webhook             string
	ReconnectAcceptList []string `default:""`
	NewBlockList        []string `default:""`
}

func main() {

	var cfg config
	err := envconfig.Process("dh", &cfg)
	if err != nil {
		slog.Error("Failed to parse env", "error", err)
		return
	}

	slog.Info("Starting dhcpresence")
	devices, err := getDevices()
	if err != nil {
		slog.Error("Failed to get cookie", "error", err)
		return
	}

	for _, device := range devices {
		if !device.Active {
			slog.Info("Device found", "device", device.Name, "active", device.Active, "lastSeen", device.LastConnection.In(time.Local))
		}
	}
}
