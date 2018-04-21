package main

import (
	"testing"
)

func TestConversation_AddNewMessage(t *testing.T) {
	// type fields struct {
	// 	msgListHead    *MessageMeta
	// 	msgListLast    *MessageMeta
	// 	cacheSize      int
	// 	maxCache       int
	// 	ContributorIDs []string
	// 	CID            string
	// }
	// type args struct {
	// 	msg *Message
	// }
	// tests := []struct {
	// 	name       string
	// 	fields     fields
	// 	args       args
	// 	wantOldMsg *Message
	// }{
	// // TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		conv := &Conversation{
	// 			msgListHead:    tt.fields.msgListHead,
	// 			msgListLast:    tt.fields.msgListLast,
	// 			cacheSize:      tt.fields.cacheSize,
	// 			maxCache:       tt.fields.maxCache,
	// 			ContributorIDs: tt.fields.ContributorIDs,
	// 			CID:            tt.fields.CID,
	// 		}
	// 		if gotOldMsg := conv.AddNewMessage(tt.args.msg); !reflect.DeepEqual(gotOldMsg, tt.wantOldMsg) {
	// 			t.Errorf("Conversation.AddNewMessage() = %v, want %v", gotOldMsg, tt.wantOldMsg)
	// 		}
	// 	})
	// }

	// TODO: Adapt this code

	// tests := []struct {
	// 	name       string
	// 	cacheLimit int
	// 	numCIDs    int
	// }{}

	// for limit := 3; limit < 50; limit++ {
	// 	for numInserts := 0; numInserts < 100; numInserts++ {
	// 		tests = append(tests, struct {
	// 			name       string
	// 			cacheLimit int
	// 			numCIDs    int
	// 		}{fmt.Sprintf("Testing cache limit %d, inserting %d", limit, numInserts), limit, numInserts})
	// 	}
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		u := &User{
	// 			headConv:       nil,
	// 			lastConv:       nil,
	// 			convCacheSize:  0,
	// 			convCacheLimit: tt.cacheLimit,
	// 		}

	// 		for idx := 0; idx < tt.numCIDs; idx++ {
	// 			u.UpdateConversationCache(idx)
	// 			if u.headConv.CID != idx {
	// 				t.Errorf("Head is of wrong CID (%d) instead of (%d).", u.headConv.CID, idx)
	// 			}
	// 			if u.lastConv.CID != max(idx-u.convCacheLimit, 0) {
	// 				t.Errorf("Last is of wrong CID (%d) instead of (%d).", u.headConv.CID, idx-u.convCacheLimit)
	// 			}
	// 		}

	// 		ptr := u.headConv
	// 		for idx := 1; idx < u.convCacheSize; idx++ {
	// 			if ptr != u.headConv && ptr.CID != ptr.head.CID-1 {
	// 				t.Errorf("List is in wrong order; %d is ahead of %d", u.headConv.CID, u.headConv.tail.CID)
	// 			}
	// 			ptr = ptr.tail
	// 		}

	// 		if ptr != u.lastConv {
	// 			fmt.Printf("List is malformed; a full traversal reached %d instead of %d\n", ptr.CID, u.lastConv.CID)
	// 		}

	// 	})
	// }
}
