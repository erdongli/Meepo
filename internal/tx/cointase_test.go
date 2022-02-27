package tx_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/erdongli/meepo/internal/tx"
	pb "github.com/erdongli/meepo/proto"
)

func TestIncrExtraNonce(t *testing.T) {
	tests := []struct {
		cbase *pb.Transaction
		want  *pb.Transaction
		err   bool
	}{
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 1}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			err: false,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns:   []*pb.TxIn{},
				TxOuts:  []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns:   []*pb.TxIn{},
				TxOuts:  []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns:   []*pb.TxIn{{}, {}},
				TxOuts:  []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns:   []*pb.TxIn{{}, {}},
				TxOuts:  []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: 0}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_NOP}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Op{Op: pb.Op_OP_NOP}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
		{
			cbase: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: math.MaxInt64}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			want: &pb.Transaction{
				Version: 0,
				TxIns: []*pb.TxIn{{ScriptSig: []*pb.Instruc{
					{Instruc: &pb.Instruc_Number{Number: 0}},
					{Instruc: &pb.Instruc_Number{Number: math.MaxInt64}},
				}}},
				TxOuts: []*pb.TxOut{{Amount: 50}},
			},
			err: true,
		},
	}

	for _, tc := range tests {
		err := tx.IncrExtraNonce(tc.cbase)
		if !reflect.DeepEqual(tc.want, tc.cbase) || tc.err != (err != nil) {
			t.Fatalf("expected: %v, err: %v, got: %v, err: %v", tc.want, tc.err, tc.cbase, err != nil)
		}
	}
}
