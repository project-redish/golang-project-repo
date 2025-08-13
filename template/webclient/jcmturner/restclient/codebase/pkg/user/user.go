package user

import (
	"fmt"

	"github.com/jcmturner/restclient"
	"{{ .ProjectName }}/pkg/client"
	"{{ .ProjectName }}/pkg/types"
)

// Create user by calling rest api
func Create(name string, id string, age int) {

	c := client.Basic()

	o := restclient.NewPostOperation()

	o.WithPath(types.CreateUserPath)

	var u types.CreateUser
	u.Name = name
	u.ID = id
	u.Age = age

	o.WithBodyDataStruct(&u)

	// Response struct for binding
	var r types.CreateUserResponse
	o.WithResponseTarget(&r)

	req, err := restclient.BuildRequest(c, o)

	if err != nil {
		fmt.Println("Error in building request")
	}

	httpcode, err := restclient.Send(req)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", *httpcode)
		fmt.Println("Message :", r.Message)
	}

}

// Get single user from server based on id
func Get(id string) {

	c := client.Basic()

	// Specify HTTP Operation type
	o := restclient.NewGetOperation()

	uri := fmt.Sprint(types.GetUserPath, id)

	o.WithPath(uri)

	// Response struct for binding
	var u types.GetUser
	o.WithResponseTarget(&u)
	req, err := restclient.BuildRequest(c, o)

	if err != nil {
		fmt.Println("Error in building request")
	}

	httpcode, err := restclient.Send(req)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", *httpcode)
		fmt.Println("Name :", u.Data.Name)
		fmt.Println("Id :", u.Data.ID)
		fmt.Println("Age :", u.Data.Age)
	}
}

// Put updates user based on id
func Put(name string, id string, age int) {

	c := client.Basic()

	// Specify HTTP Operation type
	o := restclient.NewPutOperation()

	uri := fmt.Sprint(types.UpdateUserPath, id)

	o.WithPath(uri)

	var u types.UpdateUser
	u.Name = name
	u.Age = age
	o.WithBodyDataStruct(&u)
	// Response struct for binding
	var r types.CreateUserResponse
	o.WithResponseTarget(&r)

	req, err := restclient.BuildRequest(c, o)

	if err != nil {
		fmt.Println("Error in building request")
	}

	httpcode, err := restclient.Send(req)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", *httpcode)
		fmt.Println("Message :", r.Message)
	}
}

// Find lists all users
func Find() {

	c := client.Basic()

	// Specify HTTP Operation type
	o := restclient.NewGetOperation()

	uri := fmt.Sprint(types.FindUserPath)

	o.WithPath(uri)

	// Response struct for binding
	var r types.ListUser
	o.WithResponseTarget(&r)

	req, err := restclient.BuildRequest(c, o)

	if err != nil {
		fmt.Println("Error in building request")
	}

	httpcode, err := restclient.Send(req)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", *httpcode)

	}
}
