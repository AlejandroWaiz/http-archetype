package router

import (
	"github.com/AlejandroWaiz/http-archetype/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller controller.IController
}

type IRouter interface {
	RegisterRoutes(s *gin.Engine)
}

func GetIRouter(controller controller.IController) IRouter {
	return &Router{controller}
}

func (r *Router) RegisterRoutes(s *gin.Engine) {}
