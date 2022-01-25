package server

import "github.com/gin-gonic/gin"

func New() *gin.Engine {

	s := gin.New()

	s.Use(gin.Logger())
	//s.Use(middleware.LoggerBody())
	s.Use(gin.Recovery())
	return s

}
