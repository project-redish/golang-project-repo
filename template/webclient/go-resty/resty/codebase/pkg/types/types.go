package types

const(
	// ServiceEndPoint represents server url with port number
	ServiceEndPoint = "http://localhost:8000"
)

var (

	// GetUserPath represents URI to get user
	GetUserPath = ServiceEndPoint + "/user/"

	// CreateUserPath represents URI to Create user
	CreateUserPath =ServiceEndPoint + "/user"

	// FindUserPath represents URI to get all users
	FindUserPath = ServiceEndPoint + "/user"

	// UpdateUserPath represents URI to  update user
	UpdateUserPath = ServiceEndPoint + "/user/:id"

	// DeleteUserPath represents URI to delete user
	DeleteUserPath = "/user/:id"

)

// User data struct
type User struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
	Age  int    `json:"Age"`
} 


// CreateUser struct to create user
type CreateUser struct {
	
		Name string `json:"Name"`
		ID   string `json:"ID"`
		Age  int    `json:"Age"`
}

// CreateUserResponse struct to get resp
type CreateUserResponse struct {	
	Message string `json:"message"`
}

// UpdateUser struct to create user
type UpdateUser struct {
	
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

//ListUser represents lists of users
type ListUser []struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
	Age  int    `json:"Age"`
}