package app

import (
	"fmt"
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
						fmt.Println("BIND TOGGLE UNSET", bind.ID())
					} else {
						fmt.Println("BIND TOGGLE SET", bind.ID())
					}

					EventStates[bind.ID()] = !EventStates[bind.ID()]
					return
				}

				if bind.Default {
					fmt.Println("BIND HOLD UNSET", bind.ID())
				} else {
					fmt.Println("BIND HOLD SET", bind.ID())
				}
			} else if e.KeyRelease() && !bind.Toggle {
				if bind.Default {
					fmt.Println("BIND HOLD SET", bind.ID())
				} else {
					fmt.Println("BIND HOLD UNSET", bind.ID())
				}
			}

			return
		}
	}
}
