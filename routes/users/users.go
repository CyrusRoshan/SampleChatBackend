package users

import (
	"net/http"

	"github.com/go-martini/martini"
)

func GetUsers(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {

	return 200, ""
}

func NewUser(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {

	return 200, ""
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {

	return 200, ""
}
