package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello() revel.Result {
	c.Response.Status = http.StatusOK
	return c.RenderJSON("Hello, Go World!")
}

func (c App) GetProducts() revel.Result {
	h := NewProductHandler()
	response := h.GetProducts(c)
	return c.RenderJSON(response)
}

func (c App) RegisterUser() revel.Result {
	h := NewUserHandler()
	response := h.RegisterUser(c)
	return c.RenderJSON(response)
}

func (c App) UserLogin() revel.Result {
	h := NewUserHandler()
	response := h.UserLogin(c)
	return c.RenderJSON(response)
}
