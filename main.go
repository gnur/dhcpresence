package main

import (
	"log/slog"
	"time"
)

func main() {

	activeDevices := make(map[string]Device)
	firstRun := true

	slog.Info("Starting dhcpresence")

	for {
		slog.Info("Getting devices")
		devices, err := getDevices()
		if err != nil {
			slog.Error("Failed to get cookie", "error", err)
			return
		}

		for _, device := range devices {
			_, wasActive := activeDevices[device.Key]
			if device.Active {
				if !wasActive && !firstRun {
					slog.Info("Device connected", "device", device.Name, "active", device.Active, "lastSeen", device.LastConnection.In(time.Local))
				}
				activeDevices[device.Key] = device
			} else if wasActive && !firstRun {
				slog.Info("Device disconnected", "device", device.Name, "active", device.Active, "lastSeen", device.LastConnection.In(time.Local))
				delete(activeDevices, device.Key)
			}
		}
		firstRun = false
		slog.Info("got devices", "devices", len(devices), "active", len(activeDevices))
		time.Sleep(30 * time.Second)
	}
}
