package miner

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
	privKey   *ecdsa.PrivateKey
}

func NewMiner(pool *transaction.Pool, validator *transaction.Validator) (*Miner, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return &Miner{
		pool:      pool,
		validator: validator,
		privKey:   privKey,
	}, nil
}

// Mine takes the block height, the previous block header, transactions to validate, and the targit difficulty as its
// input.
// It returns a block header that fulfills the Proof-of-Work requirement, or an error.
func (m *Miner) Mine(height int64, prevBlock [32]byte, bits uint32) (*pb.Block, error) {
	tbv := m.pool.CheckOut() // Transactions to be validated
	txs := make([]*pb.Transaction, 1, len(tbv)+1)

	var err error
	txs[0], err = transaction.NewCoinbase(height, 50, &m.privKey.PublicKey)
	if err != nil {
		return nil, err
	}

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
			PrevBlock:  prevBlock[:],
			MerkleRoot: merkleRoot[:],
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
