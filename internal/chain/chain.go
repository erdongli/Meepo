package chain

import (
	"fmt"

	"github.com/erdongli/pbchain/internal/crypto"
	"github.com/erdongli/pbchain/internal/merkle"
	"github.com/erdongli/pbchain/internal/miner"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
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
func (bc *BlockChain) Append(data [][]byte) error {
	prev := []byte{}
	if len(bc.blocks) > 0 {
		b, err := proto.Marshal(bc.blocks[len(bc.blocks)-1])
		if err != nil {
			return err
		}
		prev = crypto.Hash256(b)
	}
	mrkl := merkle.ComputeRoot(data)

	hdr, err := miner.Mine(prev, mrkl, bits)
	if err != nil {
		return err
	}

	bc.blocks = append(bc.blocks, &pb.Block{Header: hdr})
	fmt.Printf("[%d] ts: %d, merkle root: %x\n", len(bc.blocks), bc.blocks[len(bc.blocks)-1].Header.Timestamp, bc.blocks[len(bc.blocks)-1].Header.MerkleRoot)
	return nil
}
