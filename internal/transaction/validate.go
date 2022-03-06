package transaction

import (
	"github.com/erdongli/pbchain/internal/script"
	pb "github.com/erdongli/pbchain/proto"
)

type Validator struct {
	utxos *UTXOStorage
}

func NewValidator(utxos *UTXOStorage) *Validator {
	return &Validator{
		utxos: utxos,
	}
}

func (v *Validator) Validate(tx *pb.Transaction) bool {
	for i, txIn := range tx.TxIns {
		id := txIn.PrevOutput.Txid
		if len(id) != 32 {
			return false
		}
		txOut, ok := v.utxos.Get(*(*[32]byte)(id), txIn.PrevOutput.Index)
		if !ok {
			return false
		}
		if !script.ValidateTxIn(tx, txOut, i) {
			return false
		}
	}
	return true
}
