package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	list, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	// List method
	// pb.ExampleMethod(s, &grpc.....)

	logrus.Info("start listening grpc on port: ", helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(list); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
