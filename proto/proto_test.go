/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package proto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

// generate files needed
//go:generate protoc --go_out=plugins=grpc:. hello.proto


const (
	address     = "localhost:8081"
	defaultName = "world"
	port        = ":8081"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func TestProtoServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

func TestProtoClient(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewGreeterClient(conn)
	name := defaultName
	r, err := c.SayHello(context.Background(), &HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
