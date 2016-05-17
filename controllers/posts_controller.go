package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sunbirdframework/sunbird"
)

func init() {
	sunbird.RegisterController(PostsController{})
}

// PostsController ...
type PostsController struct {
	sunbird.BaseController
}

// Index ...
func (c PostsController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Index")
}

// Show ...
func (c PostsController) Show(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Show")
}

// New ...
func (c PostsController) New(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/New")
}

// Create ...
func (c PostsController) Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Create")
}

// Edit ...
func (c PostsController) Edit(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Edit")
}

// Update ...
func (c PostsController) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Update")
}

// Destroy ...
func (c PostsController) Destroy(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello from PostsController/Destroy")
}
