package script

import (
	pb "github.com/erdongli/pbchain/proto"
)

func ValidateTxIn(tx *pb.Transaction, txOut *pb.TxOut, txInIdx int) bool {
	txIn := tx.TxIns[txInIdx]

	// Get script pubkey
	scriptPubkey := txOut.ScriptPubkey

	script := make([]*pb.Instruc, 0, len(txIn.ScriptSig)+len(scriptPubkey))
	script = append(script, txIn.ScriptSig...)
	script = append(script, scriptPubkey...)

	stack := []*pb.Instruc{}
	for _, instruc := range script {
		switch v := instruc.Instruc.(type) {
		case *pb.Instruc_Op:
			if !ValidateOp(v, stack, scriptPubkey, tx, txInIdx) {
				return false
			}
		case *pb.Instruc_Number:
			stack = append(stack, instruc)
		case *pb.Instruc_Data:
			stack = append(stack, instruc)
		}
	}
	return true
}

func ValidateOp(op *pb.Instruc_Op, stack, scriptPubkey []*pb.Instruc, tx *pb.Transaction, txInIdx int) bool {
	switch op.Op {
	case pb.Op_OP_0:
		Op0(stack)
	case pb.Op_OP_DUP:
		if !OpDup(stack) {
			return false
		}
	case pb.Op_OP_EQUALVERIFY:
		if !OpEqualVerify(stack) {
			return false
		}
	case pb.Op_OP_HASH160:
		if !OpHash160(stack) {
			return false
		}
	case pb.Op_OP_CHECKSIG:
		if !OpCheckSig(stack, scriptPubkey, tx, txInIdx) {
			return false
		}
	}
	return len(stack) > 0 && isTrue(stack[len(stack)-1])
}

func isTrue(instruc *pb.Instruc) bool {
	switch v := instruc.Instruc.(type) {
	case *pb.Instruc_Number:
		return v.Number == 0
	case *pb.Instruc_Data:
		if len(v.Data) == 0 {
			return false
		}
		for _, b := range v.Data {
			if b != 0 {
				return true
			}
		}
		return false
	default:
		return false
	}
}
