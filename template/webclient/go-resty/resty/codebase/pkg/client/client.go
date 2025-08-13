package client

import (
	
	resty "github.com/go-resty/resty/v2"

)
// Basic create basic rest client
func Basic() *resty.Client {

	client := resty.New()

	return client 
}
