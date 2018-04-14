package main

import (
	"log"
	"net"

	convpb "./protobufs/dist/conversation"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func (s *Server) SortCIDs(context.Context, *convpb.CIDs) (*convpb.ConvViewsList, error) {
	ret := &convpb.ConvViewsList{}

	log.Printf("Sorting CIDs...\n")

	return ret, nil
}

func (s *Server) NewConversation(context.Context, *convpb.UIDs) (*convpb.Empty, error) {
	var ret *convpb.Empty
	return ret, nil
}

func (s *Server) NewMessage(context.Context, *convpb.Message) (*convpb.Empty, error) {
	var ret *convpb.Empty
	return ret, nil
}

func (s *Server) DeleteConversation(context.Context, *convpb.CID) (*convpb.Empty, error) {
	var ret *convpb.Empty
	return ret, nil
}

func (s *Server) DeleteHistoryOf(context.Context, *convpb.UID) (*convpb.Empty, error) {
	var ret *convpb.Empty
	return ret, nil
}

func main() {
	var srv Server

	addr := "127.0.0.1:5000"

	lis, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	sGRPC := grpc.NewServer()

	convpb.RegisterConversationManagerServer(sGRPC, &srv)
	log.Printf("Listening on %s\n", addr)

	reflection.Register(sGRPC)
	if err := sGRPC.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err.Error())
	}
}
