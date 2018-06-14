package main

import (
	"log"
	"net"

	core "github.com/itnilesh/blockchain/salpe/com/core"
	proto "github.com/itnilesh/blockchain/salpe/com/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Unable to listen at tcp port: %d and details are %v", 8080, err)
	}

	log.Print("server started..")

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{core.NewBlockChain()})
	srv.Serve(listener)

}

// Server test
type Server struct {
	Blockchain *core.BlockChain
}

// AddBlock test
func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	newBLock := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockResponse{Hash: newBLock.Hash}, nil
}

// GetBlockChain test
func (s *Server) GetBlockChain(ctx context.Context, out *proto.GetBlockChainRequest) (*proto.GetBlockChainResponse, error) {

	resp := proto.GetBlockChainResponse{}
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{Data: b.Data, Hash: b.Hash, PreviousHash: b.PrevBlockHash})

	}
	return &resp, nil
}
