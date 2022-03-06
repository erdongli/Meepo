package script

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/erdongli/pbchain/internal/crypto"
	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

func pubKeyToInstrucHash160(pk *ecdsa.PublicKey) (*pb.Instruc, error) {
	b, err := proto.Marshal(&pb.PublicKey{
		X: pk.X.Bytes(),
		Y: pk.Y.Bytes(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Instruc{Instruc: &pb.Instruc_Data{Data: crypto.Hash160(b)}}, nil
}

func instrucToPubKey(instruc *pb.Instruc) (*ecdsa.PublicKey, error) {
	switch v := instruc.Instruc.(type) {
	case *pb.Instruc_Data:
		var pk *pb.PublicKey
		err := proto.Unmarshal(v.Data, pk)
		if err != nil {
			return nil, err
		}

		return &ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     big.NewInt(0).SetBytes(pk.X),
			Y:     big.NewInt(0).SetBytes(pk.Y),
		}, nil
	default:
		return nil, fmt.Errorf("unexpected instruction type")
	}
}

func sigToInstruc(r, s *big.Int) (*pb.Instruc, error) {
	b, err := proto.Marshal(&pb.Signature{
		R: r.Bytes(),
		S: s.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.Instruc{Instruc: &pb.Instruc_Data{Data: b}}, nil
}

func instrucToSig(instruc *pb.Instruc) (*big.Int, *big.Int, error) {
	switch v := instruc.Instruc.(type) {
	case *pb.Instruc_Data:
		var sig *pb.Signature
		err := proto.Unmarshal(v.Data, sig)
		if err != nil {
			return nil, nil, err
		}

		return big.NewInt(0).SetBytes(sig.R), big.NewInt(0).SetBytes(sig.S), nil
	default:
		return nil, nil, fmt.Errorf("unexpected instruction type")
	}
}
