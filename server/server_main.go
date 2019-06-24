package server

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "perScoreAuth/perScoreProto/user"

	"google.golang.org/grpc"
)

// StartServer ...
func StartServer() {
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &Server{})
	fmt.Println("perScoreAuth server started on", os.Getenv("PORT"))
	s.Serve(lis)
}
