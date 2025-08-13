package main

import (
	"{{ .ProjectName }}/router"
	"github.com/gin-gonic/gin"
	{{ if .Logging.ImportPath }}
	_"{{ .ProjectName }}/logger"
	"{{ .Logging.ImportPath }}"
	{{end}}
	
)

func main() {

	{{ .Logging.Messages.Info }}
	{{ .Logging.Messages.Error }}
	{{ .Logging.Messages.Warn }}
	{{ .Logging.Messages.Debug }}
	
	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
