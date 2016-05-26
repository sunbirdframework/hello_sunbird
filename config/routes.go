package config

import (
	"github.com/sunbirdframework/hello_sunbird/controllers"
	"github.com/sunbirdframework/sunbird"
)

func init() {
	sunbird.Resources(controllers.PostsController{})
	sunbird.Resources(controllers.CommentsController{})
	sunbird.Get("/custom", controllers.CommentsController{}.Custom)
	// sunbird.Get("/users/:id", controllers.PostsController{}.Show)
}
