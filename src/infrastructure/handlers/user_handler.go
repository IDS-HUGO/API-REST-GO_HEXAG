package handlers

import (
	"API_REST/src/application"
	"encoding/json"
	"net/http"
	"strconv"
)

var userService = application.UserService{}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user map[string]interface{}
		json.NewDecoder(r.Body).Decode(&user)
		userService.CreateUser(user["username"].(string), user["email"].(string))
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		users, _ := userService.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	case http.MethodPut:
		var user map[string]interface{}
		json.NewDecoder(r.Body).Decode(&user)

		idFloat, ok := user["id"].(float64)
		if !ok {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		id := int(idFloat)

		userService.UpdateUser(int32(id), user["username"].(string), user["email"].(string))
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		userService.DeleteUser(int32(id))
		w.WriteHeader(http.StatusNoContent)
	}
}
