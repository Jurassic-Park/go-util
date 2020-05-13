package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type Server struct {
	CertPemPath string
	CertKeyPath string
}

func (o *Server) Get() *grpc.Server {
	// grpc server
	creds, err := credentials.NewServerTLSFromFile(o.CertPemPath, o.CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	// interceptors
	streamInterceptors, unaryInterceptors := GetInterceptors()

	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(creds), streamInterceptors, unaryInterceptors)
	grpcServer := grpc.NewServer(opts...)

	return grpcServer
}
