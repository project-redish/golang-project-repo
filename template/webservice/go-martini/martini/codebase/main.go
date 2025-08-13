package main

import (
	"{{ .ProjectName }}/router"
	"github.com/go-martini/martini"
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
	
	m := martini.Classic()
	router.RegisterRoutes(m)
	
}
