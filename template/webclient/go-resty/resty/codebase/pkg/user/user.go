package user

import (
	"fmt"

	"{{ .ProjectName }}/pkg/client"
	"{{ .ProjectName }}/pkg/types"
)

// Create user by calling rest api
func Create(name string, id string, age int) {

	c := client.Basic()

	var u types.CreateUser
	u.Name = name
	u.ID = id
	u.Age = age

	var r types.CreateUserResponse
	resp, err := c.R().
		SetBody(&u).
		SetResult(&r).
		Post(types.CreateUserPath)
	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", resp.StatusCode())
		fmt.Println("Message :", r.Message)
	}

}

// Get single user from server based on id
func Get(id string) {

	c := client.Basic()

	url := fmt.Sprint(types.GetUserPath, id)

	var u types.User

	resp, err := c.R().
		SetResult(&u).
		Get(url)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", resp.StatusCode())
		fmt.Println("Name :", u.Name)
		fmt.Println("Id :", u.ID)
		fmt.Println("Age :", u.Age)
	}
}

// Put updates user based on id
func Put(name string, id string, age int) {

	c := client.Basic()
	uri := fmt.Sprint(types.UpdateUserPath, id)

	var u types.UpdateUser
	u.Name = name
	u.Age = age

	// Response struct for binding
	var r types.CreateUserResponse

	resp, err := c.R().
		SetBody(&u).
		SetResult(&r).
		Put(uri)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", resp.StatusCode())
		fmt.Println("Message :", r.Message)
	}
}

// Find lists all users
func Find() {

	c := client.Basic()
	uri := fmt.Sprint(types.FindUserPath)

	// Response struct for binding
	var r types.ListUser

	resp, err := c.R().
		SetResult(&r).
		Get(uri)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", resp.StatusCode())

	}
}
