package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"{{ .ProjectName }}/pkg/types"
)

// Create user by calling rest api
func Create(name string, id string, age int) {
	var u types.CreateUser
	u.Name = name
	u.ID = id
	u.Age = age

	jsonValue, _ := json.Marshal(u)
	resp, err := http.Post(types.CreateUserPath, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		fmt.Println("Status Code :", resp.StatusCode)
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

}

// Get single user from server based on id
func Get(id string) {

	url := fmt.Sprint(types.GetUserPath, id)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		u, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Status Code :", resp.StatusCode)
		fmt.Println(string(u))
	}
}

// Put updates user based on id
func Put(name string, id string, age int) {

	uri := fmt.Sprint(types.UpdateUserPath, id)

	var u types.UpdateUser
	u.Name = name
	u.Age = age

	jsonValue, _ := json.Marshal(u)
	resp, err := http.NewRequest("PUT", uri, ioutil.NopCloser(strings.NewReader(string(jsonValue))))

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}

// Find lists all users
func Find() {
	resp, err := http.Get(types.FindUserPath)

	if err != nil {
		fmt.Println("Error in response", err)

	} else {
		u, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Status Code :", resp.StatusCode)
		fmt.Println(string(u))
	}
}
