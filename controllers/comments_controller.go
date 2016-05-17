package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sunbirdframework/sunbird"
)

func init() {
	sunbird.RegisterController(CommentsController{})
}

// CommentsController ...
type CommentsController struct {
	sunbird.BaseController
}

// Name ...
func (c CommentsController) Name() string {
	return "comments"
}

// Index ...
func (c CommentsController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from CommentsController/Index")
}

// Custom ...
func (c CommentsController) Custom(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from CommentsController/Custom")
}
