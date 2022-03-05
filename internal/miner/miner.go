package miner

import (
	"math"
	"time"

	"github.com/erdongli/pbchain/internal/crypto"
	"github.com/erdongli/pbchain/internal/merkle"
	"github.com/erdongli/pbchain/internal/pow"
	"github.com/erdongli/pbchain/internal/transaction"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

var TimeNow = time.Now

type Miner struct {
	pool      *transaction.Pool
	validator *transaction.Validator
}

func NewMiner(pool *transaction.Pool, validator *transaction.Validator) *Miner {
	return &Miner{
		pool:      pool,
		validator: validator,
	}
}

// Mine takes the block height, the previous block header, transactions to validate, and the targit difficulty as its
// input.
// It returns a block header that fulfills the Proof-of-Work requirement, or an error.
func (m *Miner) Mine(height int64, prevBlock []byte, bits uint32) (*pb.Block, error) {
	tbv := m.pool.CheckOut() // Transactions to be validated
	txs := make([]*pb.Transaction, 1, len(tbv)+1)
	txs[0] = transaction.NewCoinbase(height, 50)
	for _, tx := range tbv {
		if m.validator.Validate(tx) {
			txs = append(txs, tx)
		}
	}

	for {
		merkleRoot, err := merkle.ComputeRoot(txs)
		if err != nil {
			return nil, err
		}

		hdr := &pb.BlockHeader{
			Version:    0,
			PrevBlock:  prevBlock,
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
				return &pb.Block{
					Header: hdr,
					Txs:    txs,
				}, nil
			}

			// Break out upon integer overflow.
			if hdr.Nonce == math.MaxUint32 {
				break
			}

			hdr.Timestamp = uint32(TimeNow().Unix())
			hdr.Nonce++
		}

		// Try to increment extra nonce.
		if err := transaction.IncrExtraNonce(txs[0]); err != nil {
			return nil, err
		}
	}
}
