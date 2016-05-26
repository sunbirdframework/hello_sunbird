package app

import (
	"os"

	"github.com/sunbirdframework/sunbird"
)

// HelloSunbird defined your main application
var HelloSunbird *sunbird.Application

func init() {
	HelloSunbird = sunbird.App
}

// Fly ...
func Fly() {
	sunbird.InitCommands(HelloSunbird)

	if err := sunbird.App.Cmd.Execute(); err != nil {
		HelloSunbird.Router.Logger().Error(err)
		os.Exit(-1)
	}
}
