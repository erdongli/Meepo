package main

import (
	"github.com/erdongli/pbchain/internal/chain"
	"github.com/erdongli/pbchain/internal/miner"
	"github.com/erdongli/pbchain/internal/node"
	"github.com/erdongli/pbchain/internal/transaction"
)

func main() {
	storage := transaction.NewStorage()
	validator := transaction.NewValidator(storage)
	pool := transaction.NewPool()
	miner := miner.NewMiner(pool, validator)
	bchain := chain.NewBlockChain()
	node := node.NewNode(bchain, miner)
	panic(node.Run())
}
