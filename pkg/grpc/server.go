// pkg/grpc/server.go
package grpc

import (
	"net"

	"google.golang.org/grpc"
)

func StartServer(port string, register func(*grpc.Server)) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	register(s)
	return s.Serve(lis)
}
