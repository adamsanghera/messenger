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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = convManagerClient.SortCIDs(ctx, &convpb.CIDs{})
	if err != nil {
		panic(err)
	}

	log.Printf("Got the reply!!\n")
}
