package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/asdine/storm"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Device struct {
	IP       string    `storm:"index"`
	MAC      string    `storm:"id"`
	Name     string    `storm:"index"`
	LastSeen time.Time `storm:"index"`
	Expires  time.Time `storm:"index"`
	Present  bool      `storm:"index"`
}

type config struct {
	LeaseTime           time.Duration `default:"10m"`
	Database            string        `default:"./dhpresence.db"`
	Webhook             string
	ReconnectAcceptList []string `default:""`
	NewBlockList        []string `default:""`
}

const (
	EVENT_NEW_DEVICE = "network.device.new"
	EVENT_RECONNECT  = "network.device.reconnect"
	EVENT_DISCONNECT = "network.device.disconnect"
)

func main() {
	logger := logrus.WithField("source", "dhcpresence")

	var cfg config
	err := envconfig.Process("dh", &cfg)
	if err != nil {
		logger.WithError(err).Fatal("Failed to parse env")
	}

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
	flag.Parse()
	if flag.NArg() != 4 {
		logger.Fatal("not enough arguments provided")
	}

	action := flag.Arg(0)
	ip := flag.Arg(1)
	mac := flag.Arg(2)
	name := flag.Arg(3)

	if action == "commit" {
		var dev Device
		err = db.One("MAC", mac, &dev)
		sendNewEvent := false
		wasPresent := false
		if err == nil {
			wasPresent = dev.Present
		} else if err == storm.ErrNotFound {
			sendNewEvent = true
		} else if err != nil {
			logger.WithError(err).Fatal("Database error")
		}

		dev = Device{
			IP:       ip,
			MAC:      mac,
			Name:     name,
			LastSeen: time.Now(),
			Expires:  time.Now().Add(cfg.LeaseTime),
			Present:  true,
		}

		db.Save(&dev)

		eventtype := ""

		if sendNewEvent {
			logger.WithFields(logrus.Fields{
				"ip":   ip,
				"name": name,
			}).Info("saving new device")
			eventtype = EVENT_NEW_DEVICE
		} else if wasPresent {
			logger.WithFields(logrus.Fields{
				"ip":   ip,
				"name": name,
			}).Info("updating lease")
		} else {
			logger.WithFields(logrus.Fields{
				"ip":   ip,
				"name": name,
			}).Info("Welcome back device!")
			eventtype = EVENT_RECONNECT
		}

		if eventtype != "" && cfg.Webhook != "" {
			logger.Info("Triggering webhook")
			err = triggerWebhook(cfg.Webhook, eventtype, dev, cfg)
			if err != nil {
				logger.WithError(err).WithField("url", cfg.Webhook).Error("Triggering webhook failed")
			}
		}
	}
	if action == "expire" {
		var dev Device
		err = db.One("MAC", mac, &dev)

		if err != nil && err != storm.ErrNotFound {
			logger.WithError(err).Fatal("Database error")
		} else if err == storm.ErrNotFound {
			logger.Info("Device expired that we never knew")
			return
		}
		//if we get here the device was present
		if !dev.Present {
			logger.Info("Device expired that was already gone by our metrics")
			return
		}

		logger.WithFields(logrus.Fields{
			"ip":   ip,
			"name": name,
		}).Info("Goodbye device!")
		dev.Present = false
		err = db.Save(&dev)
		if err != nil {
			logger.WithError(err).Error("Failed to save to the database")
			return
		}
		logger.Info("Triggering webhook")
		err = triggerWebhook(cfg.Webhook, EVENT_DISCONNECT, dev, cfg)
		if err != nil {
			logger.WithError(err).WithField("url", cfg.Webhook).Error("Triggering webhook failed")
		}

	}
}

func triggerWebhook(hookurl, t string, dev Device, cfg config) error {

	if t != EVENT_NEW_DEVICE && len(cfg.ReconnectAcceptList) > 0 && !contains(cfg.ReconnectAcceptList, dev.Name) {
		return fmt.Errorf("Device '%s' is blocked from sending dis/reconnect events because it is not in the acceptlist", dev.Name)
	}

	if t == EVENT_NEW_DEVICE && len(cfg.NewBlockList) > 0 && contains(cfg.NewBlockList, dev.Name) {
		return fmt.Errorf("Device '%s' is blocked by the blocklist from sending new device events", dev.Name)
	}

	js, err := json.Marshal(dev)
	if err != nil {
		return err
	}

	u, err := url.Parse(hookurl)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Add("type", t)
	u.RawQuery = q.Encode()

	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(js))

	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}
	return errors.New("Webhook failed with statuscode " + resp.Status)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
