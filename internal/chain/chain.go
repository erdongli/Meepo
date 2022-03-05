package chain

import (
	pb "github.com/erdongli/pbchain/proto"
)

type BlockChain struct {
	blocks []*pb.Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*pb.Block{},
	}
}

func (bchain *BlockChain) Height() int64 {
	return int64(len(bchain.blocks))
}

func (bchain *BlockChain) GetLast() *pb.Block {
	if bchain.Height() > 0 {
		return bchain.blocks[bchain.Height()-1]
	}
	return nil
}

func (bchain *BlockChain) Append(b *pb.Block) {
	bchain.blocks = append(bchain.blocks, b)
}
