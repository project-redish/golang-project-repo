package controllers

import (
	"{{ .ProjectName }}/forms"
	"{{ .ProjectName }}/models"
	"net/http"
	"encoding/json"
	"github.com/go-martini/martini"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

var userModel = new(models.UserModel)

// UserController : struct for user controller
type UserController struct{}


// Create : function to create user 
func (user *UserController) Create(w http.ResponseWriter, r *http.Request)  {
	var data forms.CreateUserCommand
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		{{ .Logging.Messages.Error }}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = userModel.Create(data)
	if err != nil {
		{{ .Logging.Messages.Error }}
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not created"		
		json.NewEncoder(w).Encode(errorMsg)
	}else {
		var msg forms.ResponseMsg
		msg.Message = "User created"		
		json.NewEncoder(w).Encode(msg)
		
	}
}

// Get : funcation to get user
func (user *UserController) Get(w http.ResponseWriter, r *http.Request,params martini.Params) {
	id := params["id"]
	profile, err := userModel.Get(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		{{ .Logging.Messages.Error }}
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not found"		
		json.NewEncoder(w).Encode(errorMsg)
	} else {		
		json.NewEncoder(w).Encode(profile)
	}
}


// Find : List all users
func (user *UserController) Find(w http.ResponseWriter, r *http.Request) {
	list, err := userModel.Find()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		{{ .Logging.Messages.Error }}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		json.NewEncoder(w).Encode(list)
	}
}

// Update : Function to update users based on ID
func (user *UserController) Update(w http.ResponseWriter, r *http.Request,params martini.Params) {
	id := params["id"]
	data := forms.UpdateUserCommand{}
	err := json.NewDecoder(r.Body).Decode(&data)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		{{ .Logging.Messages.Error }}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userModel.Update(id, data)
	if err != nil {
		{{ .Logging.Messages.Error }}
		http.Error(w, "User not updated", http.StatusBadRequest)
		return
	} else {
		var msg forms.ResponseMsg
		msg.Message = "User Updated"		
		json.NewEncoder(w).Encode(msg)
	}
}
// Delete : Controller function to delete user based on ID
func (user *UserController) Delete(w http.ResponseWriter, r *http.Request,params martini.Params) {
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	err := userModel.Delete(id)
	if err != nil {
		{{ .Logging.Messages.Error }}
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not found"		
		json.NewEncoder(w).Encode(errorMsg)
	}else {
		var msg forms.ResponseMsg
		msg.Message = "User deleted"		
		json.NewEncoder(w).Encode(msg)
	}
}
