package main

import (
	"fmt"
	"testing"
)

func TestUser_UpdateConversationCache(t *testing.T) {
	type args struct {
		CID int
	}

	tests := []struct {
		name       string
		cacheLimit int
		numCIDs    int
	}{}

	for limit := 3; limit < 50; limit++ {
		for numInserts := 0; numInserts < 100; numInserts++ {
			tests = append(tests, struct {
				name       string
				cacheLimit int
				numCIDs    int
			}{fmt.Sprintf("Testing cache limit %d, inserting %d", limit, numInserts), limit, numInserts})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				headConv:       nil,
				lastConv:       nil,
				convCacheSize:  0,
				convCacheLimit: tt.cacheLimit,
			}

			for idx := 0; idx < tt.numCIDs; idx++ {
				u.UpdateConversationCache(idx)
				if u.headConv.CID != idx {
					t.Errorf("Head is of wrong CID (%d) instead of (%d).", u.headConv.CID, idx)
				}
				if u.lastConv.CID != max(idx-u.convCacheLimit, 0) {
					t.Errorf("Last is of wrong CID (%d) instead of (%d).", u.headConv.CID, idx-u.convCacheLimit)
				}
			}

			ptr := u.headConv
			for idx := 1; idx < u.convCacheSize; idx++ {
				if ptr != u.headConv && ptr.CID != ptr.head.CID-1 {
					t.Errorf("List is in wrong order; %d is ahead of %d", u.headConv.CID, u.headConv.tail.CID)
				}
				ptr = ptr.tail
			}

			if ptr != u.lastConv {
				fmt.Printf("List is malformed; a full traversal reached %d instead of %d\n", ptr.CID, u.lastConv.CID)
			}

		})
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
