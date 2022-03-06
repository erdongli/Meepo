package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/erdongli/pbchain/internal/chain"
	"github.com/erdongli/pbchain/internal/miner"
	"github.com/erdongli/pbchain/internal/node"
	"github.com/erdongli/pbchain/internal/transaction"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	utxos := transaction.NewUTXOStorage()
	validator := transaction.NewValidator(utxos)
	pool := transaction.NewPool()
	miner, err := miner.NewMiner(pool, validator)
	if err != nil {
		log.Fatalf("failed to create minter: %v", err)
	}
	bchain := chain.NewBlockChain()
	node := node.NewNode(bchain, miner, pool, utxos)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNodeServer(grpcServer, node)
	go grpcServer.Serve(lis)

	log.Fatalf("failed to run node: %v", node.Run())
}
