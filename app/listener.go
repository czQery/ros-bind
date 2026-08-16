package app

import (
	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func Listen(device string) {
	if device == "" {
		return
	}

	k, err := keylogger.New(device)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"device": device,
		}).Error("listener - init failed")
		return
	}
	defer k.Close()

	logrus.WithFields(logrus.Fields{
		"device": device,
	}).Info("listener - listening for input")

	for e := range k.Read() {
		switch e.Type {
		case keylogger.EvKey:
			Event(&e)
		}
	}
}
