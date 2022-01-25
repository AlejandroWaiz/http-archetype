package router

import (
	"github.com/AlejandroWaiz/http-archetype/controller"
	"github.com/AlejandroWaiz/http-archetype/middleware"
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

func (r *Router) RegisterRoutes(s *gin.Engine) {

	h := s.Group("/api/v1")
	h.GET("/health", r.controller.HealthCheck)

	v1 := s.Group("/api/v1")
	v1.Use(middleware.ValidateHeaders())

	v1.GET("/http-adapter/:transactionID", r.controller.GetTransactionController)
	v1.POST("/http-adapter", r.controller.PostTransactionController)

}
