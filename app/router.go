package app

import (
	"crypto/tls"

	"github.com/go-routeros/routeros/v3"
	"github.com/sirupsen/logrus"
)

var RouterFirewall map[string]string
var Router *routeros.Client

func ConnectRouter() {
	var err error
	Router, err = routeros.DialTLS(Config.Router.IP, Config.Router.User, Config.Router.Password, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("router - connect failed")
	}

	reply, err := Router.Run("/ip/firewall/filter/print")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("router - firewall fetch failed")
	}

	// pre fill firewall rules
	RouterFirewall = make(map[string]string)
	for _, bind := range Config.Binds {
		if bind.Target == "firewall" {
			RouterFirewall[bind.Name] = ""
		}
	}

	// set the correct id for all the saved firewall rules
	for _, re := range reply.Re {
		if _, e := RouterFirewall[re.Map["comment"]]; e {
			RouterFirewall[re.Map["comment"]] = re.Map[".id"]
		}
	}

	logrus.Info("router - connected")
}

func RouterDisableFirewall(name string, disable bool) {
	action := "disable"
	if !disable {
		action = "enable"
	}

	var id string
	if v, e := RouterFirewall[name]; e {
		id = v
	}

	if id == "" {
		logrus.WithFields(logrus.Fields{
			"err": "unknown rule",
		}).Panic("router - disable firewall failed")
	}

	_, err := Router.Run("/ip/firewall/filter/"+action, "=numbers="+id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":     err,
			"id":      id,
			"disable": disable,
		}).Error("router - disable firewall failed")
	}
}
