package http

import (
	"github.com/gin-gonic/gin"
	"simple-http-probe/config"
	"simple-http-probe/probe"
)

func StartGin(c *config.Config)  {
	r:=gin.Default()
	Routes(r)
	r.Run(c.HttpListenAddr)

}

func Routes(r *gin.Engine)  {
	api:=r.Group("/api")
	api.GET("/probe/http",probe.HttpProbe)
	api.GET("/v1", func(context *gin.Context) {
		context.String(200,"hello,我是http probe")
	})

}
