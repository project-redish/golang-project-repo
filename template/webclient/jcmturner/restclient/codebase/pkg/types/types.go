package types

const(
	// ServiceEndPoint represents server url with port number
	ServiceEndPoint = "http://localhost:8081"
)

var (

	// GetUserPath represents URI to get user
	GetUserPath = "/v1/user/"

	// CreateUserPath represents URI to Create user
	CreateUserPath ="/v1/user"

	// FindUserPath represents URI to get all users
	FindUserPath = "/v1/user"

	// UpdateUserPath represents URI to  update user
	UpdateUserPath = "/v1/user/:id"

	// DeleteUserPath represents URI to delete user
	DeleteUserPath = "/v1/user/:id"

)

// User data struct
type User struct {
	Name string `json:"Name"`
	ID   string `json:"ID"`
	Age  int    `json:"Age"`
} 

// GetUser struct
type GetUser struct {
	Data User `json:"data"`
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
type ListUser struct {
	Data []User `json:"data"`
}