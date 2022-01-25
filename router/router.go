package router

import "github.com/AlejandroWaiz/http-archetype/controller"

type Router struct {
	controller controller.IController
}

type IRouter interface {
	RegisterRoutes()
}

func RegisterRoutes() {}
