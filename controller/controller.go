package controller

import (
	"github.com/AlejandroWaiz/http-archetype/service"
)

type Controller struct {
	Service service.IService
}

type IController interface {
	GetTransactionController()
	PostTransactionController()
}

func (c *Controller) GetTransactionController() {}

func (c *Controller) PostTransactionController() {}

func GetIController(service service.IService) IController {
	return &Controller{service}
}
