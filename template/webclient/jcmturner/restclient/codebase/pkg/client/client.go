package client

import (
	"fmt"
	"os"

	"github.com/jcmturner/restclient"
	"{{ .ProjectName }}/pkg/types"
)

// Basic create basic rest client
func Basic() *restclient.Config {

	c := restclient.NewConfig()
	c.WithEndPoint(types.ServiceEndPoint)

	if err := c.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Configuration of web service not valid: %v", err)
		os.Exit(1)
	}
	return c
}
