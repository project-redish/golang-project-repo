package main

import (
	"log"

	"{{ .ProjectName }}/cmd"
)

func init() {

	//setting flags for log level
	log.SetFlags(3)
}

func main() {

	app := cmd.NewCmd()

	cmd.Run(app)

}
