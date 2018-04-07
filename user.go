package main

type User struct {
	name string
	uid  string

	// List of conversations
	headConv       *ConversationMeta // first element of LL
	lastConv       *ConversationMeta // last element of LL
	convCacheSize  int               // current cache size
	convCacheLimit int               // max size of LL
}

// ConversationMeta is a wrapper of Conversation
// that makes a linked list out of Conversation types
type ConversationMeta struct {
	head *ConversationMeta
	tail *ConversationMeta

	CID int
}

func (u *User) UpdateConversationCache(CID int) error {

	if u.convCacheSize == 0 {
		u.headConv = &ConversationMeta{
			head: nil,
			tail: nil,
			CID:  CID,
		}

		u.convCacheSize++

		u.lastConv = u.headConv
		return nil
	}

	if u.convCacheSize > u.convCacheLimit {
		// Pop the end off
		// Prepend new id
		// Pop
		penultimate := u.lastConv.head // second to last element
		penultimate.tail = nil         // popping last element (making penultimate ultimate)
		u.lastConv = penultimate       // Setting new last element

		// Prepend
		first := &ConversationMeta{ // make a new LL node
			head: nil,
			tail: u.headConv,
			CID:  CID,
		}
		u.headConv.head = first // set head of old first to be new first
		u.headConv = first      // set first to be first
		return nil
	}

	// Prepend
	first := u.headConv
	first.head = &ConversationMeta{ // make a new LL node
		head: nil,
		tail: first,
		CID:  CID,
	}
	u.headConv = first.head
	// Prepend
	// that's it

	u.convCacheSize++

	return nil
}
