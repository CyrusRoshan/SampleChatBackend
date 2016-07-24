package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"

	"github.com/cyrusroshan/SampleChatBackend/routes/chat"
	"github.com/cyrusroshan/SampleChatBackend/routes/users"
	"github.com/cyrusroshan/SampleChatBackend/store"
)

func main() {
	store.Setup()

	m := martini.New()
	router := martini.NewRouter()

	router.NotFound(func() (int, []byte) {
		return 404, []byte("Requested page not found.")
	})

	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Use(gzip.All())

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	router.Group("/chats", func(r martini.Router) {
		r.Get("/userchats/:userid", chat.GetChats)       // Get list of all chats for a user
		r.Get("/messages/:chatid", chat.ViewMessages)    // View chat messages for a chat
		r.Post("/sendmessage/:chatid", chat.SendMessage) // Send message to a chat as a user
	})

	router.Group("/users", func(r martini.Router) {
		r.Get("/get", users.GetUsers)               // Get list of all users
		r.Get("/new", users.NewUser)                // Make a new user
		r.Post("/delete/:userid", users.DeleteUser) // Delete a user with a specific id
	})

	m.MapTo(router, (*martini.Routes)(nil))
	m.Action(router.Handle)

	m.Run()
}
