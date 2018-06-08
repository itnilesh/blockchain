package main

import (
	"log"
	"net"

	proto "github.com/itnilesh/blockchain/salpe/com/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Unable to listen at tcp port: %d and details are %v", 8080, err)
	}

	log.Print("server started..")

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)

}

// Server test
type Server struct{}

// AddBlock test
func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	return new(proto.AddBlockResponse), nil
}

// GetBlockChain test
func (s *Server) GetBlockChain(context.Context, *proto.GetBlockChainRequest) (*proto.GetBlockChainResponse, error) {
	return new(proto.GetBlockChainResponse), nil
}
