package app

import (
	"github.com/sunbirdframework/sunbird"
)

// HelloSunbird defined your main application
var HelloSunbird *sunbird.Application

func init() {
	HelloSunbird = sunbird.App
}
