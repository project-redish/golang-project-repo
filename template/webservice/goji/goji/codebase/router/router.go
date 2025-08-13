package router

import (
	"{{ .ProjectName }}/controllers"
	"goji.io"
	"net/http"
	"goji.io/pat"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	
)


// RegisterRoutes creates router and routes requests
func RegisterRoutes(mux *goji.Mux) {

	{{ .Logging.Messages.Info }}
	user := new(controllers.UserController)

	mux.HandleFunc(pat.Post("/user"), user.Create)
	mux.HandleFunc(pat.Get("/user/:id"), user.Get)
	mux.HandleFunc(pat.Get("/user"), user.Find)
	mux.HandleFunc(pat.Put("/user/:id"), user.Update)
	mux.HandleFunc(pat.Delete("/user/:id"), user.Delete)

	http.ListenAndServe("localhost:8000", mux)

}
