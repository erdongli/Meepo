package main

import (
	"log"

	"github.com/erdongli/pbchain/internal/chain"
	"github.com/erdongli/pbchain/internal/miner"
	"github.com/erdongli/pbchain/internal/node"
	"github.com/erdongli/pbchain/internal/transaction"
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
	node := node.NewNode(bchain, miner, utxos)
	log.Fatalf("failed to run node: %v", node.Run())
}
