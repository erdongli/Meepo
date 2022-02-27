package miner

import (
	"math"
	"time"

	"github.com/erdongli/pbchain/internal/crypto"
	"github.com/erdongli/pbchain/internal/merkle"
	"github.com/erdongli/pbchain/internal/pow"
	"github.com/erdongli/pbchain/internal/tx"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

var TimeNow = time.Now

// Mine takes the block height, the previous block header, transactions to validate, and the targit difficulty as its
// input.
// It returns a block header that fulfills the Proof-of-Work requirement, or an error.
func Mine(height int64, prevBlock *pb.BlockHeader, txs []*pb.Transaction, bits uint32) (*pb.BlockHeader, error) {
	prevBytes := []byte{}
	if height != 0 {
		var err error
		prevBytes, err = proto.Marshal(prevBlock)
		if err != nil {
			return nil, err
		}
	}

	// Create a new slice of transactions with coinbase being the first element.
	txsCb := make([]*pb.Transaction, len(txs)+1)
	txsCb[0] = tx.NewCoinbase(height, 50)
	copy(txsCb[1:], txs)

	for {
		merkleRoot, err := merkle.ComputeRoot(txsCb)
		if err != nil {
			return nil, err
		}

		hdr := &pb.BlockHeader{
			Version:    0,
			PrevBlock:  crypto.Hash256(prevBytes),
			MerkleRoot: merkleRoot,
			Timestamp:  uint32(TimeNow().Unix()),
			Bits:       bits,
			Nonce:      0,
		}

		v := pow.NewValidator(bits)
		for {
			b, err := proto.Marshal(hdr)
			if err != nil {
				return nil, err
			}

			if v.Validate(crypto.Hash256(b)) {
				return hdr, nil
			}

			// Break out upon integer overflow.
			if hdr.Nonce == math.MaxUint32 {
				break
			}

			hdr.Timestamp = uint32(TimeNow().Unix())
			hdr.Nonce++
		}

		// Try to increment extra nonce.
		if err := tx.IncrExtraNonce(txsCb[0]); err != nil {
			return nil, err
		}
	}
}
