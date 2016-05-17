package environments

import (
	"github.com/sunbirdframework/hello_sunbird/config/app"
)

func init() {
	app.HelloSunbird.Environment("test", func() {
		app.HelloSunbird.Router.Logger().Info("Initialized with env=test")
	})
}
