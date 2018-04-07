package main

type User struct {
	name string
	uid  string

	//
	headConv *ConversationMeta // first element of LL
	lastConv *ConversationMeta // last element of LL
}

// ConversationMeta is a wrapper of Conversation
// that makes a linked list out of Conversation types
type ConversationMeta struct {
	head *ConversationMeta
	tail *ConversationMeta

	CID string
}
