package script

import (
	"crypto/ecdsa"

	"github.com/erdongli/pbchain/internal/crypto"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

func Op0(stack []*pb.Instruc) {
	stack = append(stack, &pb.Instruc{Instruc: &pb.Instruc_Data{}})
}

func OpVerify(stack []*pb.Instruc) bool {
	if len(stack) == 0 {
		return false
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return isTrue(top)
}

func OpDup(stack []*pb.Instruc) bool {
	if len(stack) == 0 {
		return false
	}
	stack = append(stack, stack[len(stack)-1])
	return true
}

func OpEqual(stack []*pb.Instruc) bool {
	if len(stack) < 2 {
		return false
	}
	l, r := stack[len(stack)-1], stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	num := int64(0)
	if proto.Equal(l, r) {
		num = 1
	}
	stack = append(stack, &pb.Instruc{Instruc: &pb.Instruc_Number{Number: num}})
	return true
}

func OpEqualVerify(stack []*pb.Instruc) bool {
	if OpEqual(stack) == false {
		return false
	}
	return OpVerify(stack)
}

func OpHash160(stack []*pb.Instruc) bool {
	if len(stack) == 0 {
		return false
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	data, ok := top.Instruc.(*pb.Instruc_Data)
	if !ok {
		return false
	}
	hash := crypto.Hash160(data.Data)
	stack = append(stack, &pb.Instruc{Instruc: &pb.Instruc_Data{Data: hash}})
	return true
}

func OpCheckSig(stack, scriptPubkey []*pb.Instruc, tx *pb.Transaction, idx int) bool {
	// Get public key and signature from stack
	if len(stack) < 2 {
		return false
	}
	pk, err := toPublicKey(stack[len(stack)-1])
	if err != nil {
		return false
	}
	r, s, err := toSignature(stack[len(stack)-2])
	if err != nil {
		return false
	}
	stack = stack[:len(stack)-2]

	// Set all TxIn scripts to nil
	txCpy := proto.Clone(tx).(*pb.Transaction)
	for i, txIn := range txCpy.TxIns {
		txIn.ScriptSig = nil
		if i == idx {
			// Set TxIns[txInIdx]'s script sig to the previous output's script pubkey
			txIn.ScriptSig = scriptPubkey
		}
	}

	// Marshall the transaction copy
	b, err := proto.Marshal(txCpy)
	if err != nil {
		return false
	}

	return ecdsa.Verify(pk, crypto.Hash256(b), r, s)
}
