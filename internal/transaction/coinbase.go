package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"math"

	"github.com/erdongli/pbchain/internal/script"
	pb "github.com/erdongli/pbchain/proto"
)

// NewCoinbase creates a new coinbase transaction.
// A coinbase transaction contains a single input, whose previous output is set to nil, and script sig
// set to [<block height>, <extra nonce>].
func NewCoinbase(height int64, amount uint64, pk *ecdsa.PublicKey) (*pb.Transaction, error) {
	script, err := script.P2PKH(pk)
	if err != nil {
		return nil, err
	}
	return &pb.Transaction{
		Version: 0,
		TxIns: []*pb.TxIn{
			{
				ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: height}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
				},
			},
		},
		TxOuts: []*pb.TxOut{
			{
				Amount:       amount,
				ScriptPubkey: script,
			},
		},
	}, nil
}

// IncreExtraNonce increments coinbase's extra nonce.
func IncrExtraNonce(cbase *pb.Transaction) error {
	if len(cbase.TxIns) != 1 {
		return fmt.Errorf("invalid number of TxIns: %d", len(cbase.TxIns))
	}
	txIn := cbase.TxIns[0]

	// script sig should be in the form of [<block height>, <extra nonce>].
	if len(txIn.ScriptSig) != 2 {
		return fmt.Errorf("invalid scipt sig")
	}
	instruc := txIn.ScriptSig[1]

	switch instruc.Instruc.(type) {
	case *pb.Instruc_Number:
		n := instruc.GetNumber()
		// Check for integer overflow.
		if n == math.MaxInt64 {
			return fmt.Errorf("integer overflow when incrementing extra nonce")
		}
		txIn.ScriptSig[1] = &pb.Instruc{Instruc: &pb.Instruc_Number{Number: n + 1}}
		return nil
	default:
		return fmt.Errorf("invalid extra nonce")
	}
}
