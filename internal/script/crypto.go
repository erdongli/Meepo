package script

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"

	pb "github.com/erdongli/pbchain/proto"
	"google.golang.org/protobuf/proto"
)

func toPublicKey(instruc *pb.Instruc) (*ecdsa.PublicKey, error) {
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

func toSignature(instruc *pb.Instruc) (*big.Int, *big.Int, error) {
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
