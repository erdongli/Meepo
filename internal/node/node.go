package node

import (
	"fmt"

	"github.com/erdongli/pbchain/internal/chain"
	"github.com/erdongli/pbchain/internal/crypto"
	"github.com/erdongli/pbchain/internal/miner"
	"google.golang.org/protobuf/proto"
)

const bits = uint32(20)

type Node struct {
	bchain *chain.BlockChain
	miner  *miner.Miner
}

func NewNode(bchain *chain.BlockChain, miner *miner.Miner) *Node {
	return &Node{
		bchain: bchain,
		miner:  miner,
	}
}

// Append is a naive method just so that the block chain can be assembled.
func (n *Node) Run() error {
	for {
		height := n.bchain.Height()
		lastBytes := []byte{}
		if last := n.bchain.GetLast(); last != nil {
			var err error
			lastBytes, err = proto.Marshal(last.Header)
			if err != nil {
				return err
			}
			lastBytes = crypto.Hash256(lastBytes)
		}

		block, err := n.miner.Mine(height, lastBytes, bits)
		if err != nil {
			return err
		}

		n.bchain.Append(block)
		fmt.Printf("[%d] ts: %d, merkle root: %x\n", height, block.Header.Timestamp, block.Header.MerkleRoot)
	}
}
