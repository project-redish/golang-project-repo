package main

import (
	"os"

	"{{ .ProjectName }}/cmd"
)

func main() {
	cmd.Commands(os.Args[1:])
}
