package chain

import (
	"fmt"

	"github.com/erdongli/pbchain/internal/miner"
	pb "github.com/erdongli/pbchain/proto"
)

const bits = uint32(20)

type BlockChain struct {
	blocks []*pb.Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*pb.Block{},
	}
}

// Append is a naive method just so that the block chain can be assembled.
func (bc *BlockChain) Append(txs []*pb.Transaction) error {
	height := len(bc.blocks)
	var prev *pb.BlockHeader
	if height > 0 {
		prev = bc.blocks[height-1].Header
	}

	hdr, err := miner.Mine(int64(height), prev, txs, bits)
	if err != nil {
		return err
	}

	bc.blocks = append(bc.blocks, &pb.Block{Header: hdr})
	fmt.Printf("[%d] ts: %d, merkle root: %x\n", len(bc.blocks), bc.blocks[len(bc.blocks)-1].Header.Timestamp, bc.blocks[len(bc.blocks)-1].Header.MerkleRoot)
	return nil
}
