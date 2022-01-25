package controller

import (
	"fmt"
	"net/http"

	"github.com/AlejandroWaiz/http-archetype/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.IService
}

type IController interface {
	GetTransactionController(ctx *gin.Context)
	PostTransactionController(ctx *gin.Context)
	HealthCheck(ctx *gin.Context)
}

func (c *Controller) GetTransactionController(ctx *gin.Context) {

	transactionID := ctx.Param("transactionID")

	headers := getHeadersMap(ctx.Request.Header)

	stc, body, err := c.Service.GetTransactionService(headers, transactionID)

	if err != nil {
		fmt.Printf(`Error getting data from datasource adapter: %s`+"\n", err)
		ctx.Data(http.StatusServiceUnavailable, "application/json; charset=utf-8", body)
		return
	}

	ctx.Data(stc, "application/json; charset=utf-8", body)

}

func (c *Controller) PostTransactionController(ctx *gin.Context) {

	headers := getHeadersMap(ctx.Request.Header)

	ctxBody := ctx.Request.Body

	stc, body, err := c.Service.PostTransactionService(headers, ctxBody)

	if err != nil {
		ctx.Data(http.StatusServiceUnavailable, "application/json; charset=utf-8", body)
		return
	}

	ctx.Data(stc, "application/json; charset=utf-8", body)

}

//HealthCheck actuator
func (ctr *Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

func GetIController(service service.IService) IController {
	return &Controller{service}
}

func getHeadersMap(h http.Header) map[string]string {
	m := make(map[string]string)
	for k, v := range h {
		m[k] = v[0]
	}
	return m
}
