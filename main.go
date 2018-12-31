package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Trendydata/fileserver/internal"
	"github.com/Trendydata/fileserver/pkg"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pkg.RegisterFileServer(grpcServer, &internal.FileServer{})
	println("Ready to recieve chunks!")
	grpcServer.Serve(lis)
}
