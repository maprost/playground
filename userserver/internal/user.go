package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/maprost/playground/userserver/client"
	"github.com/maprost/playground/userserver/internal/data"
	"github.com/maprost/playground/userserver/internal/util"
)

func initUser(router *mux.Router) {
	router.Path("/user").Methods(http.MethodPost).HandlerFunc(createUser)
	router.Path("/user/{email}").Methods(http.MethodDelete).HandlerFunc(deleteUser)
	router.Path("/user").Methods(http.MethodGet).HandlerFunc(getUsers)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var userDto client.UserDto
	err := json.NewDecoder(request.Body).Decode(&userDto)
	if err != nil {
		http.Error(writer, "Can't open UserDto", http.StatusNotFound)
		return
	}

	user := data.UserAccount{
		Email:        userDto.Email,
		PasswordHash: userDto.PasswordHash,
	}

	err = user.Insert()
	if err != nil {
		http.Error(writer, "Can't store user in database.", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	email, ok := mux.Vars(request)["email"]
	if !ok {
		http.Error(writer, "Can't delete user out of database.", http.StatusBadRequest)
		return
	}

	err := data.DeleteUser(email)
	if err != nil {
		http.Error(writer, "Can't delete user out of database.", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func getUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := data.GetUsers()
	if err != nil {
		http.Error(writer, "Can't load user out of database.", http.StatusBadRequest)
		return
	}

	// convert to dto
	userDtos := util.ConvertUsersToUserDtos(users)

	js, err := json.Marshal(userDtos)
	if err != nil {
		http.Error(writer, "Some internal problem", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(js)
}
