package app

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var Config config

type config struct {
	Keyboard string         `json:"keyboard"`
	Mouse    string         `json:"mouse"`
	Mikrotik configMikrotik `json:"mikrotik"`
	Binds    []configBind   `json:"binds"`
}

type configMikrotik struct {
	IP       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type configBind struct {
	Key     string `json:"key"`
	Toggle  bool   `json:"toggle"`
	Default bool   `json:"default"`
	Target  string `json:"target"`
	Name    string `json:"name"`
}

func (bind *configBind) ID() string {
	return bind.Key + ":" + bind.Name
}

func LoadConfig() {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("config - read failed")
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("config - unmarshal failed")
	}

	EventMutex.Lock()
	EventStates = make(map[string]bool)
	for _, bind := range Config.Binds {
		EventStates[bind.ID()] = bind.Default
	}
	EventMutex.Unlock()

	logrus.Info("config - loaded")
}
