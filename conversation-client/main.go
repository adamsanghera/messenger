package main

import (
	"context"
	"log"
	"time"

	convpb "github.com/adamsanghera/messenger/protobufs/dist/conversation"
	"google.golang.org/grpc"
)

func main() {
	convManagerAddr := "172.104.239.167:5000"

	conn, err := grpc.Dial(convManagerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	convManagerClient := convpb.NewConversationManagerClient(conn)

	sum := int64(0)
	count := int64(0)
	for {
		count++
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		start := time.Now()
		_, err = convManagerClient.SortCIDs(ctx, &convpb.CIDs{})
		if err != nil {
			panic(err)
		}
		rt := time.Since(start).Nanoseconds()
		sum += rt
		log.Printf("Round trip in %d nanoseconds", rt)
		log.Printf("Average round trip has been %d nanoseconds", sum/count)
	}
}
