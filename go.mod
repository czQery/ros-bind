module github.com/czQery/ros-bind

go 1.26.5

require (
	github.com/MarinX/keylogger v0.0.0-20240620105846-48ca9d01f566
	github.com/sirupsen/logrus v1.10.0
)

require golang.org/x/sys v0.13.0 // indirect

replace github.com/MarinX/keylogger => github.com/czQery/keylogger v0.0.0-20260816124620-5873603078ef
