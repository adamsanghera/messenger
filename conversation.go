package main

import "time"

// Message is the guts of a message
type Message struct {
	author    string
	timestamp time.Time
	content   string
}

// MessageMeta is a linked list node wrapper around the Message object
type MessageMeta struct {
	// LL Component
	head *MessageMeta
	tail *MessageMeta

	// Actual message
	msg *Message
}

// Conversation is a linked list of messages, with attending metadata
type Conversation struct {
	msgListHead *MessageMeta // head element of LL
	msgListLast *MessageMeta // last element of LL
	cacheSize   int
	maxCache    int

	ContributorIDs []string // UIDs of users in conversation
	CID            string   // unique ID of convo
}

// AddNewMessage prepends to the LL iff cache size is less than the cap.
// Otherwise, it prepends to the LL and kicks off the last.
func (conv *Conversation) AddNewMessage(msg *Message) (oldMsg *Message) {

	if conv.cacheSize < conv.maxCache {
		oldMsg = nil

		// Prepend
		first := conv.msgListHead
		first.head = &MessageMeta{ // make a new LL node
			head: nil,
			tail: first,
			msg:  msg,
		}
		conv.msgListHead = first.head

		conv.cacheSize++
	} else {
		oldMsg = conv.msgListLast.msg

		// Pop
		penultimate := conv.msgListLast.head // second to last element
		penultimate.tail = nil               // popping last element (making penultimate ultimate)
		conv.msgListLast = penultimate       // Setting new last element

		// Prepend
		first := &MessageMeta{ // make a new LL node
			head: nil,
			tail: conv.msgListHead,
			msg:  msg,
		}
		conv.msgListHead.head = first // set head of old first to be new first
		conv.msgListHead = first      // set first to be first
	}
	return
}
