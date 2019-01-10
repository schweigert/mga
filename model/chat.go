package model

// Chat is a message
type Chat struct {
	Character Character
	Line      string
}

// ChatSet is a set of chats
type ChatSet struct {
	Set []Chat
}
