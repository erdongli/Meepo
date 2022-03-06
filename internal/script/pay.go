package script

import (
	"crypto/ecdsa"

	pb "github.com/erdongli/pbchain/proto"
)

func P2PKH(pk *ecdsa.PublicKey) ([]*pb.Instruc, error) {
	h160, err := pubKeyToInstrucHash160(pk)
	if err != nil {
		return nil, err
	}
	return []*pb.Instruc{
		{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_DUP}},
		{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_HASH160}},
		h160,
		{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_EQUALVERIFY}},
		{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_CHECKSIG}},
	}, nil
}
