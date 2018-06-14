package main

import (
	"flag"
	"log"

	proto "github.com/itnilesh/blockchain/salpe/com/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("Add", false, "Add new block")
	listFlag := flag.Bool("List", false, "List blocks")
	flag.Parse()

	host := "localhost:8080"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	client = proto.NewBlockchainClient(conn)

	if err != nil {
		log.Fatalf("Can not connect to server located at %s", host)
	}

	if *addFlag {
		addBlock()

	} else if *listFlag {
		listBlocks()
	} else {
		log.Print("Invalid Request")
	}

}

func addBlock() {
	req := proto.AddBlockRequest{Data: "Nilesh"}
	resp, err := client.AddBlock(context.Background(), &req)

	if err == nil {
		log.Printf("New block added %s ", resp.Hash)
	} else {
		log.Fatal("Last operation failed to execute")
	}
}

func listBlocks() {
	req := proto.GetBlockChainRequest{}
	resp, err := client.GetBlockChain(context.Background(), &req)

	log.Printf("size %d", len(resp.Blocks))
	if err == nil {
		for _, b := range resp.Blocks {
			log.Printf(" ==> %s ", b.Data)
		}
	} else {
		log.Fatal("Last operation failed to execute")
	}
}
