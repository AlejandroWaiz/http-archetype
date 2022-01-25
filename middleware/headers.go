package middleware

import (
	"fmt"
	"net/http"

	"github.com/AlejandroWaiz/http-archetype/response"
	"github.com/gin-gonic/gin"
)

//ValidateHeaders validate corp headers
func ValidateHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := []string{"X-Txref", "X-Cmref", "X-Rhsref", "X-Chref", "X-Country", "X-Commerce", "X-Api-Key", "Entityid"}
		for _, v := range h {
			if _, m := ctx.Request.Header[v]; m != true {
				fmt.Println(v)

				ctx.AbortWithStatusJSON(http.StatusBadRequest, response.MessageDescription{Message: response.Message{
					Code:        http.StatusBadRequest,
					Description: "The request header is not valid",
				}})
				break
			}
		}
	}
}
