package node

import (
	"context"
	"fmt"

	"github.com/erdongli/pbchain/internal/chain"
	"github.com/erdongli/pbchain/internal/crypto"
	"github.com/erdongli/pbchain/internal/miner"
	"github.com/erdongli/pbchain/internal/transaction"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

const bits = uint32(25)

type Node struct {
	pb.UnimplementedNodeServer

	bchain *chain.BlockChain
	miner  *miner.Miner
	pool   *transaction.Pool
	uxtos  *transaction.UTXOStorage
}

func NewNode(bchain *chain.BlockChain, miner *miner.Miner, pool *transaction.Pool, uxtos *transaction.UTXOStorage) *Node {
	return &Node{
		bchain: bchain,
		miner:  miner,
		pool:   pool,
		uxtos:  uxtos,
	}
}

// Append is a naive method just so that the block chain can be assembled.
func (n *Node) Run() error {
	for {
		height := n.bchain.Height()
		prevBlock := [32]byte{}
		if prev := n.bchain.GetLast(); prev != nil {
			var err error
			b, err := proto.Marshal(prev.Header)
			if err != nil {
				return err
			}
			prevBlock = crypto.Hash256(b)
		}

		block, err := n.miner.Mine(height, prevBlock, bits)
		if err != nil {
			return err
		}

		n.bchain.Append(block)
		n.uxtos.Update(block)
		fmt.Printf("[%d] ts: %d, merkle root: %x\n", height, block.Header.Timestamp, block.Header.MerkleRoot)
		fmt.Printf("UTXOs: %v\n", n.uxtos)
	}
}

func (n *Node) NewTx(_ context.Context, req *pb.NewTxRequest) (*pb.NewTxResponse, error) {
	n.pool.CheckIn(req.Tx)
	return &pb.NewTxResponse{}, nil
}
