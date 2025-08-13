package main

import (
	"fmt"

	"{{ .ProjectName }}/pkg/user"
)

func main() {

	fmt.Println("********Create User Call*****")
	user.Create("user1", "1", 30)
	fmt.Println("********Completed************")
	fmt.Println("")
	fmt.Println("********Get User Call*****")
	user.Get("1")
	fmt.Println("********Completed*****")
	fmt.Println("********Update User Call*****")
	user.Put("user2", "1", 33)
	fmt.Println("********Completed*****")
	fmt.Println("********List User Call*****")
	user.Find()
	fmt.Println("********Completed*****")

}
