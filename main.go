package main

import (
	"github.com/czQery/ros-bind/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("main - ros-bind by Štěpán Aubrecht")
	app.LoadConfig()

	go app.Listen(app.Config.Keyboard)
	go app.Listen(app.Config.Mouse)

	select {}
}
