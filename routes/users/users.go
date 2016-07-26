package users

import (
	"net/http"
	"strconv"

	"github.com/cyrusroshan/API/utils"
	"github.com/cyrusroshan/SampleChatBackend/store"
	"github.com/go-martini/martini"
)

func GetUsers(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	return 200, string(utils.MustMarshal(store.Users))
}

func NewUser(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	userName := params["userName"]
	userId := store.NewUser(userName)
	return 200, string(utils.MustMarshal(userId))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	useridString := params["userid"]
	userid, err := strconv.Atoi(useridString)
	utils.PanicIf(err)

	delete(store.Users, userid)
	return 200, ""
}
