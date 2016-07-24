package store

import (
	"errors"
	"math/rand"
	"time"
)

type User struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Chats UserChats `json:"-"`
}

type Message struct {
	Content  string `json:"content"`
	FromUser int    `json:"fromuser"`
	Time     int64  `json:"time"`
}

type Chat struct {
	Id       int       `json:"id"`
	Users    []int     `json:"users"`
	LastSent int64     `json:"lastsent"`
	Messages []Message `json:"-"`
}

type UserChats map[int]*Chat

var chatIDstore map[int]bool
var Users map[int]User
var ChatStore UserChats
var RandomGenerator *rand.Rand

func Setup() {
	ChatStore = make(UserChats)
	Users = make(map[int]User)
	randSource := rand.NewSource(time.Now().UnixNano())
	RandomGenerator = rand.New(randSource)
}

func NewChat(users []int) int {
	var chatID int
	for {
		chatID = RandomGenerator.Int()
		if chatIDstore[chatID] == false {
			chatIDstore[chatID] = true
			break
		}
	}

	chat := Chat{
		Id:       chatID,
		Users:    users,
		LastSent: time.Now().Unix(),
		Messages: make([]Message, 0),
	}

	for _, user := range users {
		Users[user].Chats[chatID] = &chat
	}

	ChatStore[chatID] = &chat
	return chatID
}

func (user User) NewMessage(content string, chatId int) error {
	newMessage := Message{
		Content:  content,
		FromUser: user.Id,
		Time:     time.Now().Unix(),
	}

	if user.Chats[chatId].Messages == nil {
		return errors.New("Chat id doesn't exist")
	}

	user.Chats[chatId].LastSent = time.Now().Unix()
	user.Chats[chatId].Messages = append(user.Chats[chatId].Messages, newMessage)

	return nil
}
