package router

import (
	"{{ .ProjectName }}/controllers"
	"github.com/go-martini/martini"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

// RegisterRoutes creates router and routes requests
func RegisterRoutes(m *martini.ClassicMartini) {
	
	{{ .Logging.Messages.Info }}
	m.Group("/v1", func(r martini.Router) {
		user := new(controllers.UserController)
		r.Post("/user",  user.Create)
		r.Get("/user/:id", user.Get)
		r.Get("/user", user.Find)
		r.Put("/user/:id", user.Update)
		r.Delete("/user/:id", user.Delete)
	})

	m.Run()
}
