package main

import (
	"fmt"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening at :12345")
	fmt.Println(router)
	{{.Logging.Messages.Info}}
	//http.ListenAndServe(":8080", router)
}
