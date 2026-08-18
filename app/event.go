package app

import (
	"sync"

	"github.com/MarinX/keylogger"
)

var EventMutex sync.Mutex
var EventStates map[string]bool

func Event(e *keylogger.InputEvent) {
	for _, bind := range Config.Binds {
		if bind.Key == e.KeyString() {

			if e.KeyPress() {
				if bind.Toggle {
					EventMutex.Lock()
					defer EventMutex.Unlock()

					if EventStates[bind.ID()] {
						go EventRun(bind, true)
					} else {
						go EventRun(bind, false)
					}

					EventStates[bind.ID()] = !EventStates[bind.ID()]
					return
				}

				if bind.Default {
					go EventRun(bind, true)
				} else {
					go EventRun(bind, false)
				}
			} else if e.KeyRelease() && !bind.Toggle {
				if bind.Default {
					go EventRun(bind, false)
				} else {
					go EventRun(bind, true)
				}
			}

			return
		}
	}
}

func EventRun(bind configBind, disable bool) {
	switch bind.Target {
	case "firewall":
		RouterDisableFirewall(bind.Name, disable)
	}
}
