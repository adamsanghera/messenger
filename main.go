package main

import (
	"net/http"
)

// CID -> Conversation
var convLedger map[string]*Conversation

// UID -> User
var userLedger map[string]*User

var updateWorker DBUpdateWorker

func HandleNewMessage(w http.ResponseWriter, req *http.Request) {
	// auth, err := req.Cookie("author-id")
	// if err != nil {
	// 	panic(err)
	// }
	// newMsg := &Message{
	// 	author:    req.Cookie("author-id"),
	// 	timestamp: time.Now(),
	// }

	// triggerFunc(req.Form.Get("Message-Content"), req.Cookie("CID"))
}

// newMsg comes from the HTTP Req;
// CID comes from the HTTP Req;
// dbuw
func triggerFunc(newMsg Message, CID string) {
	old := convLedger[CID].AddNewMessage(&newMsg)
	updateWorker.NewMessage(*old)
}

func main() {
	http.HandleFunc("/newMessage", HandleNewMessage)
}
