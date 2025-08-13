package models

import (
	"{{ .ProjectName }}/forms"
	"fmt"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

// User : Struct for getting user data
type User struct {
	Name   string        `json:name`
	ID   string        `json:id`
	Age float32       `json:age`
	
}
//UserModel : sruct
type UserModel struct{}

// Create : A model function to create user
func (m *UserModel) Create(data forms.CreateUserCommand) error {
	{{ .Logging.Messages.Info }}
	var err error
	fmt.Println("User created with below data")
	fmt.Println("Name : ",data.Name)
	fmt.Println("ID : ",data.ID)
	fmt.Println("Age : ",data.Age)
	err = nil
	return err
}

// Get : Function to retrive user base on id 
func (m *UserModel) Get(id string) (user User, err error) {
	{{ .Logging.Messages.Info }}
	user = User{Name: "user"+id, ID: id, Age: 30}
	return user, err
}
// Find : Function to retrive list of all users
func (m *UserModel) Find() (list []User, err error) {
	{{ .Logging.Messages.Info }}
	user1 := User{Name: "user1", ID: "id1", Age: 30}
	list = append(list,user1)
	user2 := User{Name: "user2", ID: "id2", Age: 25}
	list = append(list,user2)
	return list, err
}

// Update : Function to update user details based on id recived from user
func (m *UserModel) Update(id string, data forms.UpdateUserCommand) (err error) {
	{{ .Logging.Messages.Info }}
	// Update logic goes here
	fmt.Println("User with id "+ id +" is updated with below data")
	fmt.Println("Name : ",data.Name)
	fmt.Println("Age : ",data.Age)
	err = nil
	return err
}

// Delete : Function to delete user based on id
func (m *UserModel) Delete(id string) (err error) {
	{{ .Logging.Messages.Info }}
	fmt.Println("User with id "+ id +" is deleted successfully")
	err = nil
	return err
}
