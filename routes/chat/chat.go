package chat

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"

	"github.com/cyrusroshan/SampleChatBackend/store"
	"github.com/cyrusroshan/SampleChatBackend/utils"
)

func GetChats(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	useridString := params["userid"]
	userid, err := strconv.Atoi(useridString)
	utils.PanicIf(err)

	user := store.Users[userid]
	return 200, string(utils.MustMarshal(user.Chats))
}

func ViewMessages(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	chatidString := params["chatid"]
	chatid, err := strconv.Atoi(chatidString)
	utils.PanicIf(err)

	chat := store.ChatStore[chatid]
	return 200, string(utils.MustMarshal(chat.Messages))
}

func SendMessage(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	chatidString := params["chatid"]
	chatid, err := strconv.Atoi(chatidString)
	utils.PanicIf(err)

	decoder := json.NewDecoder(r.Body)
	var message store.Message
	err = decoder.Decode(&message)
	utils.PanicIf(err)

	store.ChatStore[chatid].Messages = append(store.ChatStore[chatid].Messages, message)

	return 200, ""
}
