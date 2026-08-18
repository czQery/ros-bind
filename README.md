# ros-bind

- Example config.json

```json
{
	"keyboard": "/dev/input/by-id/usb-HP__Inc_HyperX_Alloy_Origins_65-event-kbd",
	"mouse": "/dev/input/by-id/usb-HP__Inc_HyperX_Pulsefire_Haste_2_Wireless-event-mouse",
	"router": {
		"ip": "192.168.1.1:8729",
		"user": "ros-bind",
		"password": "loooong_password"
	},
	"binds": [
		{
			"key": "L_ALT",
			"toggle": false,
			"default": false,
			"target": "firewall",
			"name": "ros-bind drop web"
		}
	]
}

```
