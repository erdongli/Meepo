package main

import (
	"log"

	"github.com/erdongli/pbchain/internal/chain"
	pb "github.com/erdongli/pbchain/proto"
)

func main() {
	bc := chain.NewBlockChain()
	for {
		if err := bc.Append([]*pb.Transaction{}); err != nil {
			log.Fatalf("failed to append to block chain: %v", err)
		}
	}
}
